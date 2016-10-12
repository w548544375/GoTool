package ssession

import (
	"smessage"
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
	uuID            int
			     //关联的socket
	client          SSocket
			     //发送列表
	messages        *MessageList
			     //close回掉函数
	OnClose         func()

	endSync         sync.WaitGroup

	needNotifyWrite bool // 是否需要通知写线程关闭
}

func (self *SSession) Send(message *smessage.SMessage) {
	if message != nil {
		self.messages.Push(message)
	}
}

func (self *SSession) UuID() {
	return self.uuID
}

func (self *SSession) sendTread() {
	willExit := false
	for !willExit {
		msg := self.messages.Pick()
		if msg != nil {
			willExit = self.client.SendMessage(msg)
		}
	}
	self.needNotifyWrite = false
	// 通知发送线程ok
	self.endSync.Done()
	self.client.Close()
}

func (self *SSession) recvThread(evq SEventQueue) {
	willExit := false
	for !willExit {
		msg, err := self.client.RecvMessage()
		if err != nil {
			willExit = true
		}
		evq.PostMessage(msg, self)
	}

	if self.needNotifyWrite {
		self.client.Close()
	}
	// 通知发送线程ok
	self.endSync.Done()
}

//新建session
func newSession(socket SSocket, evq SEventQueue) *SSession {
	self := &SSession{
		client:          socket,
		needNotifyWrite: true,
		messages:NewMessageList(),
	}
	// 布置接收和发送2个任务
	self.endSync.Add(2)

	go func() {
		self.endSync.Wait()
		if self.OnClose != nil {
			self.OnClose()
		}
	}()

	go self.recvThread(evq)

	go self.sendTread()

	return self
}