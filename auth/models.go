package auth


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
	Name         string         `orm:"size(80)"`
	Permissions  []*Permission  `orm:"rel(m2m)"`

}

func (group *Group) TableName() string {
	return "auth_group"
}

func (group *Group) TableUnique() [][]string {

	 return [][]string{
	 	[]string{"Name"},
	 }
}

/*
A mixin class that adds the fields and methods necessary to support
     Group and Permission model using the ModelBackend.
 */
type User struct {
	Username        string
	Email           string

    IsSuperuser     bool            `orm:"default(false)"`

    Groups          []*Group        `orm:"rel(m2m)"`
	UserPermissions []*Permission   `orm:"rel(m2m)"`
}

