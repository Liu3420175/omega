package sessions

import "time"

type Session struct {
	SessionKey     string         `orm:"pk"`
	SessionDdata   string
	ExpireDate     time.Time
}


func (session *Session) TableName() string  {
	return "auth_session"

}
