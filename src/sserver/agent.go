package main

import (
	"ssession"
	"ssocket"
	"sync"
	"smessage"
	"sbuffer"
)

func main() {
	var wait sync.WaitGroup
	acc := ssocket.NewSocketAcceptor(&wait)
	acc.Evq.RegisterEventHandler(ssession.SessionAccepted, func(a interface{}) {
		event := a.(*ssession.SEvent)
		message := smessage.NewSMessage(0,0,len(ssession.Meet),nil,nil,sbuffer.Wrap(ssession.Meet))
		event.Shooter().Send(message)
	})
	acc.Start(":2287")
	//遍历事件
	go func() {
		for  {
		  event := acc.Evq.PopEvent()
		  if nil != event{
			  acc.Evq.RunEvent(event)
		  }
		}
	}()

	wait.Wait()

}
