package ssession

import (
	"smessage"
	"sync"
)

type MessageList struct {
	//
	list []*smessage.SMessage

	guard sync.Mutex

	listCond *sync.Cond
}

//放入消息
func (self *MessageList) Push(msg *smessage.SMessage) {
	self.guard.Lock()
	defer self.guard.Unlock()
	self.list = append(self.list, msg)
	self.listCond.Signal()
}

//取消息
func (self *MessageList) Pick() *smessage.SMessage {
	self.guard.Lock()
	defer self.guard.Unlock()
	if len(self.list) == 0 {
		return nil
	}
	message := self.list[0]
	self.list = self.list[1:]
	return message
}

func (self *MessageList) Reset() {
	self.guard.Lock()
	defer self.guard.Unlock()
	self.list = self.list[0:0]
}

func NewMessageList() *MessageList {
	self := &MessageList{
		list: make([]*smessage.SMessage, 0),
	}
	self.listCond = sync.NewCond(&self.guard)
	return self
}
