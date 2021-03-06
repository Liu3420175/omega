package auth

import (
	"strings"
	"github.com/astaxie/beego/orm"
	"time"
	"errors"
	"reflect"
	"omega/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
The permissions system provides a way to assign permissions to specific
    users and groups of users.
 */
type Permission struct {
	Id           int64          `orm:"pk;auto"`
    Name         string         `orm:"size(255)"`
	Content_type *ContentType   `orm:"rel(fk);null;column(content_type);on_delete(set_null)"`
	Codename      string        `orm:"size(100)"`
	Groups       []*Group `orm:"reverse(many)"`
}

func (permission *Permission) TableName() string  {
	return "auth_permission"

}

func (permission *Permission) TableUnique() [][]string {
	return [][]string{
		[]string{"Name","Content_type"},
	}
}


func (persission *Permission) natural_key() []string {
	keys := persission.Content_type.natural_key()
	return []string{keys[0],keys[1],persission.Codename}
}


func (persission *Permission) String() string{
	return "Permission:" + persission.Name
}


type Group struct {
	Id           int64          `orm:"pk;auto"`
	Name         string         `orm:"size(80);unique"`
	Note         string         `orm:"size(225)"`
	Permissions  []*Permission  `orm:"rel(m2m)"`
	Users        []*User        `orm:"reverse(many)"`

}

func (group *Group) TableName() string {
	return "auth_group"
}

func (group *Group) String() string {
	return "Group:" + group.Name
}


type Session struct {
	SessionKey     string         `orm:"pk;size(64);null"`
	SessionData    string         `orm:"size(552);null"`
	ExpireDate     time.Time      `orm:"type(datetime)"`
	UserId         int64          `orm:"null"`
}


func (session *Session) TableName() string  {
	return "auth_session"

}

func (session *Session) String() string {
	return "Session:" + session.SessionKey
}


type User struct {
	Id int64 					    `orm:"pk;auto"`
	UserName        string          `orm:"sieze(150);unique"`
	Email           string          `orm:"size(128);unique"`
	FirstName       string          `orm:"size(30)"`
	LastName        string          `orm:"size(30)"`
	Password        string          `orm:"size(128)"`
	Phone           string          `orm:"size(20)"`
	LastLogin       time.Time       `orm:"auto_now_add;type(datetime)"`
	DateJoined      time.Time       `orm:"auto_now_add;type(datetime)"`
	IsActive        bool            `orm:"default(true)"`
	IsStaff         bool            `orm:"default(false)"`
	IsAdmin         bool            `orm:"default(false)"`
	IsSuperuser     bool            `orm:"default(false)"`
	Content_type    *ContentType    `orm:"rel(fk);null;column(content_type);on_delete(set_null)"`
    Groups          []*Group        `orm:"rel(m2m)"`
}


func (user *User) String() string {
	return "User:" + user.Email
}


func (user *User) TableName() string {
	return "user_user"
}



func (user *User) GetUsername() string{
	return user.UserName
}


func (user *User) IsAnonymous() bool {
	return false
}


func CompareUser(user1 *User,user2 *User) bool{
	v1 := reflect.ValueOf(user1)
	v2 := reflect.ValueOf(user2)
	return v1.Interface() == v2.Interface()
}


func (user *User) IsAuthenticated() bool {
	return true
}


func (user *User) SetPassword(raw_password string )  {
	// TODO
	//o := orm.NewOrm()
	user.Password = MakePassword(raw_password,"")
	//_,err := o.Update(user)
	//if err == nil {
	//	return true }else{
	//	return false
	//}

}


func (user *User) CheckPassword(raw_password string) bool{
	return CheckPassword(raw_password,user.Password)
}

func (user *User) GetFullName()  string {

	full_name := strings.Join([]string{user.FirstName,user.LastName}," ")
	return strings.Trim(full_name," ")
}

func (user *User) GerShortName() string {
	return user.FirstName
}


func (user *User) GetGroupPermissions() orm.ParamsList{

	o := orm.NewOrm()
    group_name := new(Group)
    permission_name := new(Permission)
    var group_ids,permissions orm.ParamsList
    o.QueryTable(group_name).Filter("Users__Id",user.Id).ValuesFlat(&group_ids,"id")
    o.QueryTable(permission_name).Filter("Groups__Id__in",group_ids).ValuesFlat(&permissions,"codename")
	return permissions
}



func (user *User) GetSessionAuthHash() string{
	key_salt := "omega.auth.models.User.get_session_auth_hash"
	return SaltedHhmac(key_salt,user.Password,"")
}



func (user *User)GetPerm(perm string) bool{
    perms := user.GetGroupPermissions()
    for _,v := range perms{
    	if v == perm{
    		return true
		}
	}
	return false
}



func create_user(fields map[string]string) User {
	var user User
    username := fields["UserName"]
    if len(username) == 0{
    	panic("The given username must be set")
	}
	email := fields["Email"]
	if len(email) == 0{
		panic("The given email must be set")
	}
    password := fields["Password"]
    user.Password = MakePassword(password,"")
    user.UserName = username
    user.Email = email
    user.FirstName = fields["FirstName"]
    user.LastName = fields["LastName"]
    user.Phone = fields["Phone"]
    user.DateJoined = time.Now()
    user.IsActive = true

    return user
}




func CreateUser(fields map[string]string) (*User,error ){
	user := create_user(fields)
	user.IsStaff = true
	user.IsSuperuser = false
	user.IsAdmin = true
	o := orm.NewOrm()
    _,err := o.Insert(&user)
    if err == nil {
    	return &user,nil
	}else{
		// TODO
		return nil,errors.New("Create Error" + err.Error())
	}
}



func CreateSuperuser (fields map[string]string) (*User,error ) {
	user := create_user(fields)
	user.IsStaff = true
	user.IsSuperuser = true
	user.IsAdmin = true
	o := orm.NewOrm()
	_,err := o.Insert(&user)
	if err == nil {
		return &user,nil
	}else{
		// TODO
		return nil,errors.New("Create Error" + err.Error())
	}
}


func init() {

	databases := conf.DATABASES
	database := databases["default"]
	fmt.Print(database)
	orm.RegisterDriver("mysql",orm.DRMySQL)
	connent := database["USER"] + ":" + database["PASSWORD"] + "@tcp(" + database["HOST"] + ":" + database["PORT"] + ")/" + database["NAME"] + "?charset=utf8"
	orm.RegisterDataBase("","mysql",connent)
	orm.RegisterModel(
		new(Group),
		new(Permission),
		new(ContentType),
		new(Session),
		new(User),
	)
}