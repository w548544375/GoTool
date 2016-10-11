package SMessage

import (
	"sbuffer"
)


//用于发送的消息结构体
type SMessage struct {
	headData    SBuffer.SBuffer
	headLength  int
	extraData   SBuffer.SBuffer
	extraLength int
	mainData    SBuffer.SBuffer
	mainLength  int
}

func (msg *SMessage) HeadLength() int{
	return msg.headLength
}

func (msg *SMessage) ExtraLength() int{
	return msg.extraLength
}

func (msg *SMessage) MainLength() int {
	return msg.mainLength
}

//取得buff的head
func (msg *SMessage) Head() []byte {
	return msg.headData.Bytes()
}

//ID 或者人物坐标数据
func (msg *SMessage) Extra() []byte {
	return msg.extraData.Bytes()
}

//主要消息
func (msg *SMessage) Main() []byte {
	return msg.mainData.Bytes()
}

