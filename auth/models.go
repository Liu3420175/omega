package auth

import "strings"

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




type User struct {
    AbstractBaseUser
	Content_type    *ContentType       `orm:"rel(fk);null;column(content_type);on_delete(set_null)"`
    Groups          []*Group           `orm:"rel(m2m)"`
	Permissions     []*Permission      `orm:"rel(m2m)"`
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


func (user *User) IsAuthenticated() bool {
	return true
}


func (user *User) SetPassword(string) bool {
	// TODO
	return false
}


func (user *User) GetFullName()  string {

	full_name := strings.Join([]string{user.FirstName,user.LastName}," ")
	return strings.Trim(full_name," ")
}

func (user *User) GerShortName() string {
	return user.FirstName
}