package smessage

import (
	"sbuffer"
	"fmt"
)

type IMessage interface {
	HeadLength() int16
	ExtraLength() int16
	MainLength() int
	Head() *sbuffer.SBuffer
	Extra() *sbuffer.SBuffer
	Main() *sbuffer.SBuffer
	SetHead(length int16, head *sbuffer.SBuffer)
	SetExtra(length int16, extra *sbuffer.SBuffer)
	SetMain(length int, main *sbuffer.SBuffer)
}

//用于发送的消息结构体
type SMessage struct {
	headData    *sbuffer.SBuffer
	headLength  int16
	extraData   *sbuffer.SBuffer
	extraLength int16
	mainData    *sbuffer.SBuffer
	mainLength  int
}

func (self *SMessage) String() string {
	return fmt.Sprintf("head:%d\nData:%v\nExtra:%d\nData:%v\nMain:%d\nData:%v\n",
	self.headLength,self.headData.Bytes(),self.extraLength,self.extraData.Bytes(),self.mainLength,self.mainData.Bytes())
}


func (msg *SMessage) HeadLength() int16 {
	return msg.headLength
}

func (msg *SMessage) ExtraLength() int16 {
	return msg.extraLength
}

func (msg *SMessage) MainLength() int {
	return msg.mainLength
}

//取得buff的head
func (msg *SMessage) Head() *sbuffer.SBuffer {
	return msg.headData
}

//ID 或者人物坐标数据
func (msg *SMessage) Extra() *sbuffer.SBuffer {
	return msg.extraData
}

//主要消息
func (msg *SMessage) Main() *sbuffer.SBuffer {
	return msg.mainData
}

func (msg *SMessage) SetHead(length int16, head *sbuffer.SBuffer) {
	if length > 0 && head.Limit() > 0 && head.Limit() == int(length) {
		msg.headData = head
		msg.headLength = length
	}
}

func (msg *SMessage) SetExtra(length int16, extra *sbuffer.SBuffer) {
	if length > 0 && extra.Limit() > 0 && extra.Limit() == int(length) {
		msg.extraData = extra
		msg.extraLength = length
	}
}

func (msg *SMessage) SetMain(length int, main *sbuffer.SBuffer) {
	if length > 0 && main.Limit() > 0 && main.Limit() == length {
		msg.mainData = main
		msg.mainLength = length
	}
}

func NewSMessage(headlen,extralen int16,mainlen int,head,extra,main *sbuffer.SBuffer) *SMessage {
	return &SMessage{headData:head,
			 headLength:headlen,
			 extraData:extra,
			extraLength:extralen,
			mainData:main,
			mainLength:mainlen,
				}
}