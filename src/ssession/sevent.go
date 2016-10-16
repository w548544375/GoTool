package ssession

import (
	"smessage"
	"fmt"
)

type IEvent interface {
	EventID() int16
	Shooter() *SSession
}

type SEvent struct {
	evId    int16     //事件id
	shooter *SSession // 事件的发起者
	message *smessage.SMessage
}

func (self *SEvent) String() string{
	return fmt.Sprintf("{evId:%d,shooterID:%d,message:%v}\n",self.evId,self.Shooter().UuID(),self.message)
}

//获取事件ID
func (self *SEvent) EventID() int16 {
	return self.evId
}

func (self *SEvent) Shooter() *SSession {
	return self.shooter
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
