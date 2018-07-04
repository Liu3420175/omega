package auth


/*
The permissions system provides a way to assign permissions to specific
    users and groups of users.
 */
type Permission struct {
	Id           int64          `orm:"pk;auto"`

}