package sessions

import "time"

type Session struct {
	SessionKey     string         `orm:"pk"`
	SessionDdata   string
	ExpireDate     time.Time      `orm:"type(datetime)"`
}


func (session *Session) TableName() string  {
	return "auth_session"

}
