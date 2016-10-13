package smessage

import (
	"sbuffer"
)

//用于发送的消息结构体
type SMessage struct {
	headData    *SBuffer.SBuffer
	headLength  int16
	extraData   *SBuffer.SBuffer
	extraLength int16
	mainData    *SBuffer.SBuffer
	mainLength  int
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

func (msg *SMessage) SetHead(length int, head *SBuffer.SBuffer) {
	if length > 0 && head.Limit() > 0 && head.Limit() == int(length) {
		msg.headData = head
		msg.headLength = length
	}
}

func (msg *SMessage) SetExtra(length int, extra *SBuffer.SBuffer) {
	if length > 0 && extra.Limit() > 0 && extra.Limit() == int(length) {
		msg.extraData = extra
		msg.extraLength = length
	}
}

func (msg *SMessage) SetMain(length int, main *SBuffer.SBuffer) {
	if length > 0 && main.Limit() > 0 && main.Limit() == length {
		msg.mainData = main
		msg.mainLength = length
	}
}
