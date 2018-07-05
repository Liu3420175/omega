package auth

import (
	"time"
)

type AbstractBaseUser struct {
	Username        string          `orm:"sieze(150);unique"`
	FirstName       string          `orm:"size(30)"`
	LastName        string          `orm:"size(30)"`
	Email           string          `orm:"size(128),unique"`
	Password        string          `orm:"size(128)"`
	Phone           string          `orm:"size(20)"`
	LastLogin       time.Time       `orm:"auto_now_add;type(datetime)"`
	DateJoined      time.Time       `orm:"auto_now_add;type(datetime)"`
	IsActive        bool            `orm:"default(true)"`
	IsStaff         bool            `orm:"default(false)"`
	IsSuperuser     bool            `orm:"default(false)"`
}