package main

import (
	"fmt"
	"sync"
	"time"
)

// type IEvent interface{
// 	Event()
// 	Name()
// 	EventID()
// }

// type Event struct{
// 	op int32
// 	eventID int32
// 	name string

// }

type IEventQueue interface {
	AddEvent(string)
	RegisterEventHandler(string, func(*int64))
	UnRegisterEventHandler(string)
	RunEvent()
}

type EventQueue struct {
	eventList    []string
	registedFunc map[string]func(*int64)
	staus        sync.Mutex
	counter      int64
}

func (self *EventQueue) AddEvent(name string) {
	self.staus.Lock()
	defer self.staus.Unlock()
	self.eventList = append(self.eventList, name)
}

func (self *EventQueue) RegisterEventHandler(name string, handler func(*int64)) {
	if _, ok := self.registedFunc[name]; !ok {
		self.registedFunc[name] = handler
	}
}

func (self *EventQueue) RunEvent() {
	self.staus.Lock()
	defer self.staus.Unlock()
	if 0 < len(self.eventList) {
		event := self.eventList[0]
		self.eventList = self.eventList[1:]
		if fn, ok := self.registedFunc[event]; ok {
			go fn(&self.counter)
		}
	}
}

func (self *EventQueue) UnRegisterEventHandler(name string) {
	if _, ok := self.registedFunc[name]; ok {
		delete(self.registedFunc, name)
	}
}

func NewEventQueue() *EventQueue {
	return &EventQueue{
		eventList:    make([]string, 1),
		registedFunc: make(map[string]func(*int64)),
		counter:      0}
}

func main() {
	handleCount := 0
	queue := NewEventQueue()
	var wait sync.WaitGroup
	brun := true
	wait.Add(1)
	queue.RegisterEventHandler("test", Test)
	go func() {
		for {
			queue.AddEvent("test")
			//time.Sleep(time.Millisecond * 10)
		}
	}()
	for i := 0; i < 100; i++ {
		go func() {
			for brun {
				queue.RunEvent()
				handleCount++
				//time.Sleep(time.Millisecond * 10)
			}
		}()
	}
	go func() {
		time.Sleep(10000 * time.Millisecond)
		brun = false
		wait.Done()
	}()
	// wait.Wait()
	fmt.Printf("Handled Event:%d\n", handleCount)
}

func Test(times *int64) {
	//fmt.Printf("This is the InnerCall %d\n", *times)
	var temp int64 = *times
	for temp > 0 {
		temp--
	}

	*times++
}
