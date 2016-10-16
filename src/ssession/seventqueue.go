package ssession

/*
 *  事件队列，存储事件
 */
import (
	"fmt"
	"sync"
)

type IEventQueue interface {
	RegisterEventHandler(int16, func(interface{})) //注册事件处理函数
	PushEvent(*SEvent)                             //向事件队列写事件
	PopEvent() *SEvent                             //取一个事件
}

type SEventQueue struct {
	eventList []*SEvent //事件队列

	registedHandlers map[int16]func(interface{}) //事件注册的处理函数

	guard sync.Mutex //并发守护
}

//注册事件处理函数
func (self *SEventQueue) RegisterEventHandler(eventID int16, handler func(interface{})) {
	if _, ok := self.registedHandlers[eventID]; ok {
		delete(self.registedHandlers, eventID)
	}
	self.registedHandlers[eventID] = handler
}

//写入事件
func (self *SEventQueue) PushEvent(event *SEvent) {
	fmt.Printf("事件到达:%v\n", event)
	self.guard.Lock()
	defer self.guard.Unlock()
	self.eventList = append(self.eventList, event)
}

func (self *SEventQueue) PopEvent() *SEvent {
	self.guard.Lock()
	defer self.guard.Unlock()
	//避免出现index out of range 异常
	if len(self.eventList) == 0 {
		return nil
	}
	event := self.eventList[0]          //取第一个事件
	self.eventList = self.eventList[1:] //切片
	return event
}

func (self *SEventQueue) RunEvent(event *SEvent){
	if fn,ok := self.registedHandlers[event.evId];ok{
		fn(event)
	}
}

//创建事件队列
func NewSEventQueue() *SEventQueue {
	return &SEventQueue{
		eventList:        make([]*SEvent, 0),
		registedHandlers: make(map[int16]func(interface{})),
	}
}
