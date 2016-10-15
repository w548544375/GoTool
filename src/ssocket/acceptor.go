package ssocket

import (
	"fmt"
	"net"
	"ssession"
	"sync"
)

type SocketAcceptor struct {
	running bool

	Mgr *ssession.SessionMgr

	Evq *ssession.SEventQueue

	listener net.Listener

	wait *sync.WaitGroup
}

func (self *SocketAcceptor) Start(address string) {

	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server start Listening.....")

	self.listener = ln
	self.running = true
	self.wait.Add(1)

	go func() {
		for self.running {
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			fmt.Printf("客户端连接：%v\n", conn)
			go func() {
				client := ssession.NewSession(ssession.NewSocket(conn), self.Evq)
				client.OnClose = func() {
					self.Mgr.Remove(client)
				}
				fmt.Printf("session 创建成功，%v\n", client)
				self.Mgr.Add(client)
				self.Evq.PushEvent(ssession.NewSEvent(client, ssession.SessionAccepted, nil))
			}()
		}
		self.wait.Done()
	}()

}

func (self *SocketAcceptor) Stop() {
	if !self.running {
		return
	}
	self.running = false
	self.listener.Close()
}

func NewSocketAcceptor(waitO *sync.WaitGroup) *SocketAcceptor {
	return &SocketAcceptor{
		running: true,
		Mgr:     ssession.NewSessionMgr(),
		Evq:     ssession.NewSEventQueue(),
		wait:    waitO,
	}
}
