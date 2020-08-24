package model

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type SSession struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Validator [16]byte  `json:"validator"`
	//Expire    time.Timer `json:"expire"`
}

// 全局session池
var SessionPool = sync.Map{}

const SessionPeriod = 1 * time.Hour

func SetSession(u SUser, agent string) *SSession {
	session := new(SSession)
	session.ID = uuid.New()
	session.Username = u.Username
	//session.Validator = md5.Sum([]byte(agent))
	SessionPool.Delete(u.Username)
	SessionPool.Store(u.Username, session)
	return session
}

func (s *SSession) Validate() bool {
	var _s = new(SSession)
	v, ok := SessionPool.Load(s.Username)
	if !ok {
		return false
	}
	_s, ok = v.(*SSession)
	if !ok {
		return false
	}
	if s.ID == _s.ID {
		return true
	}
	return false
}
