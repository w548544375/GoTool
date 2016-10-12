package ssocket

import (
	"net"
	"ssession"
	"fmt"
)

type SocketAcceptor struct {

	//sessionMgr

	listener net.Listener

	running bool
}

var eventQueue = &ssession.SEventQueue{}

func (self *SocketAcceptor) Start(address string){
	
	ln,err := net.Listen("tcp",address)
	if err != nil {
		panic(err)
	}
	self.listener = ln

	self.running =true

	go func(){
		for self.running{
			conn,err := ln.Accept()
			if(err != nil){
				break
			}
			fmt.Println(conn)
			//session := ssession.NewSession(conn,eventQueue)

		}
	}();
}