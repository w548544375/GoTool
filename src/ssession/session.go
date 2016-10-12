package ssession

import (
	"smessage"
	"ssocket"
	"sync"
)

type ISession interface {

	//发送封包
	Send(interface{})
	//获取session的id
	UuID()
	//关闭socket
	Close()
}

type SSession struct {
	//session ID
	UuID int
	//关联的socket
	client ssocket.SSocket
	//发送列表
	messages MessageList
	//close回掉函数
	OnClose func()

	endSync sync.WaitGroup

	needNotifyWrite bool // 是否需要通知写线程关闭
}

func (self *SSession) Send(message *smessage.SMessage) {
	if message != nil {
		self.messages.Push(message)
	}
}

