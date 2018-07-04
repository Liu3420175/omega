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
}

func (permission *Permission) TableName() string  {
	return "auth_permission"

}

func (permission *Permission) TableUnique() [][]string {
	return [][]string{
		[]string{"Name","Content_type"},
	}
}