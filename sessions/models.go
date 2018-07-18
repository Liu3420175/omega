package sessions

import "time"

type Session struct {
	SessionKey     string         `orm:"pk"`
	SessionData    string
	ExpireDate     time.Time      `orm:"type(datetime)"`
	UserId         int64
}


func (session *Session) TableName() string  {
	return "auth_session"

}

func (session *Session) String() string {
	return "Session:" + session.SessionKey
}