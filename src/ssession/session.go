package ssession

import (
	"smessage"
	"sync"
	"fmt"
)

type ISession interface {
	//发送封包
	Send(*smessage.SMessage)
	//获取session的id
	UuID() int32
	//关闭socket
	Close()
}

type SSession struct {
	//session ID
	uuID int32
	//关联的socket
	client *SSocket
	//发送列表
	messages *MessageList
	//close回掉函数
	OnClose func()

	endSync sync.WaitGroup

	needNotifyWrite bool // 是否需要通知写线程关闭
}

func (self *SSession) Close() {
	self.client.Close()
}

func (self *SSession) Send(message *smessage.SMessage) {
	if message != nil {
		self.messages.Push(message)
	}
}




func (self *SSession) UuID() int32 {
	return self.uuID
}

func (self *SSession) sendTread() {
	willExit := false
	for !willExit {
		msg := self.messages.Pick()
		if msg != nil {
			fmt.Printf("发送消息\n:%v\n",msg)
			willExit = !self.client.SendMessage(msg)
		}
	}
	fmt.Println("线程执行结束")
	self.needNotifyWrite = false
	// 通知发送线程ok
	self.endSync.Done()
	self.client.Close()
}

func (self *SSession) recvThread(evq *SEventQueue) {
	willExit := false
	for !willExit {
		buff := make([]byte,256)
		msg, err := self.client.Recv(buff)
		if err != nil {
			fmt.Printf("客户端断开连接：%v\n",self)
			willExit = true
		}
		if msg == -1{
			fmt.Print("收到EOF\n")
		}else {
			fmt.Printf("收到的消息为：%v\n", msg)
			evq.PushEvent(NewSEvent(self, msg.Main().GetShortFrom(0), msg))
		}

	}

	if self.needNotifyWrite {
		fmt.Printf("退出接受线程：%v",self.needNotifyWrite)
		self.client.Close()
	}
	// 通知发送线程ok
	self.endSync.Done()
}

//新建session
func NewSession(socket *SSocket, evq *SEventQueue) *SSession {
	self := &SSession{
		client:          socket,
		needNotifyWrite: true,
		messages:        NewMessageList(),
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
