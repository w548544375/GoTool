package ssession

import (
	"smessage"
)

type IEvent interface {
	EventID() int16
}

type SEvent struct {
	evId    int16     //事件id
	shooter *SSession // 事件的发起者
	message *smessage.SMessage
}

//获取事件ID
func (self *SEvent) EventID() int16 {
	return self.evId
}

/*定义Sevent的方法*/

//新建事件
func NewSEvent(poster *SSession, eventID int16, data *smessage.SMessage) *SEvent {
	return &SEvent{
		evId:    eventID, //事件id
		shooter: poster,  //事件的发送者
		message: data,    //事件的数据
	}
}
