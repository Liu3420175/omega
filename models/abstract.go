package models

import "time"

type Base struct {
	Id           int64          `orm:"pk;auto"`
	DateCreated  time.Time     	`orm:"auto_now_add;type(datetime)"`
	DateModified time.Time    	`orm:"auto_now;type(datetime)"`
	Creator      string         `orm:"null;size(128)"`
	Lastmodifier string         `orm:"null;size(128)"`
}
