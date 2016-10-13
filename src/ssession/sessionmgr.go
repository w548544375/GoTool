package ssession

import (
	"sync"
	"sync/atomic"
)

type ISessionMgr interface {
	Add(ISession)
	Remove(ISession)
}

type SessionMgr struct {
	// 管理的session
	managedSes map[int32]*SSession

	//并发保护
	guard sync.RWMutex

	//当前的sessionId
	idAcs int32
}

const totalTryCount = 100

func (self *SessionMgr) Add(client ISession) {
	self.guard.Lock()
	defer self.guard.Unlock()

	tryCount := totalTryCount
	var id int32
	for i := 0; i < tryCount; i++ {
		id = atomic.AddInt32(&self.idAcs, 1)
		if _, ok := self.managedSes[id]; !ok {
			sSession := client.(*SSession)
			sSession.uuID = id
			self.managedSes[id] = sSession
			break
		}
	}
}

//移除连接
func (self *SessionMgr) Remove(client ISession) {
	self.guard.Lock()
	defer self.guard.Unlock()

	if client != nil {
		delete(self.managedSes, client.UuID())
	}
}
