package auth

import (
	"strings"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/pkg/errors"
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
	Groups       []*Permission  `orm:"reverse(many)"`
	Users        []*User        `orm:"reverse(many)"`
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


type User struct {
	Id int64 					    `orm:"pk;auto"`
	Username        string          `orm:"sieze(150);unique"`
	Email           string          `orm:"size(128),unique"`
	FirstName       string          `orm:"size(30)"`
	LastName        string          `orm:"size(30)"`
	Password        string          `orm:"size(128)"`
	Phone           string          `orm:"size(20)"`
	LastLogin       time.Time       `orm:"auto_now_add;type(datetime)"`
	DateJoined      time.Time       `orm:"auto_now_add;type(datetime)"`
	IsActive        bool            `orm:"default(true)"`
	IsStaff         bool            `orm:"default(false)"`
	IsSuperuser     bool            `orm:"default(false)"`
	Content_type    *ContentType    `orm:"rel(fk);null;column(content_type);on_delete(set_null)"`
    Groups          []*Group        `orm:"rel(m2m)"`
	Permissions     []*Permission   `orm:"rel(m2m)"`
}


func (user *User) String() string {
	return "User:" + user.Email
}


func (user *User) TableName() string {
	return "user_user"
}



func (user *User) GetUsername() string{
	return user.Username
}


func (user *User) IsAnonymous() bool {
	return false
}


func CompareUser(user1 *User,user2 *User) bool{
	// TODO
	return false
}


func (user *User) IsAuthenticated() bool {
	return true
}


func (user *User) SetPassword(raw_password string ) bool {
	// TODO
	o := orm.NewOrm()
	user.Password = MakePassword(raw_password,"")
	_,err := o.Update(user)
	if err == nil {
		return true
	}else{
		return false
	}

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


func (user *User) GetGroupPermissions() []string{
	var permissions  []string

	return permissions
}

func (user *User) GetSessionAuthHash() string{
	key_salt := "omega.auth.models.User.get_session_auth_hash"
	return SaltedHhmac(key_salt,user.Password,"")
}


func create_user(fields map[string]string) User {
	var user User
    username := fields["Username"]
    if len(username) == 0{
    	panic("The given username must be set")
	}
	email := fields["Email"]
	if len(email) == 0{
		panic("The given email must be set")
	}
    password := fields["Password"]
    user.Password = MakePassword(password,"")
    user.Username = username
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
	user.IsStaff = false
	user.IsSuperuser = false
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
	o := orm.NewOrm()
	_,err := o.Insert(&user)
	if err == nil {
		return &user,nil
	}else{
		// TODO
		return nil,errors.New("Create Error" + err.Error())
	}
}


