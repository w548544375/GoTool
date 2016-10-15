package main

import (
	"fmt"
	"ssession"
	"ssocket"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	acc := ssocket.NewSocketAcceptor(&wait)
	acc.Evq.RegisterEventHandler(ssession.SessionAccepted, func(interface{}) {
		fmt.Println("Hello ,New Client Connected...")
	})
	acc.Start(":2287")

	go func() {
		for {
			event := acc.Evq.PopEvent()
			if nil == event {
				continue
			}

		}
	}()

	wait.Wait()

}
