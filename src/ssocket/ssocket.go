package ssocket

import (
	"io"
	"net"
	"sbuffer"
	"smessage"
)

const (
	DEFAULT_HEAD_LENGTH = 18
)

var LENGTHERROR = error.Error("Packet Length mismatch")
var ERR_LENGTH_EXTRA = error.Error("Extra Packet Length mismatch")
var ERR_LENGTH_MAIN = error.Error("Main Packet Length mismatch")

type SSocket struct {
	net.TCPConn
}

func (socket *SSocket) Recv(buff []byte) (int,error) {
	length := 0
	for {
		n, err := socket.Read(buff[length:])
		if n > 0 {
			length += n
		}
		if err != nil {
			if err != io.EOF {
				return 0,err
			}
			break
		}
	}
	return length,nil
}

//接受数据 组装为message
func (socket *SSocket) RecvMessage() (*smessage.SMessage,error) {
	msg := new(smessage.SMessage)
	buff := make([]byte, DEFAULT_HEAD_LENGTH)
	socket.Recv(buff)
	//验证数据包
	headBuf := SBuffer.Wrap(buff)
	msgType := headBuf.GetShortFrom(0)
	extraLength := headBuf.GetShortFrom(2)
	mainLength := headBuf.GetIntFrom(4)
	validate := headBuf.GetShortFrom(16)
	if msgType ^ extraLength ^ mainLength != validate {
		return nil,LENGTHERROR
	}
	msg.SetHead(DEFAULT_HEAD_LENGTH, headBuf)
	if extraLength != 0 {
		buff = make([]byte, extraLength)
		n,err := socket.Recv(buff)
		if err != nil {
			return nil,err
		}
		if n == extraLength {
			msg.SetExtra(extraLength, SBuffer.Wrap(buff))
		}else{
			return nil,ERR_LENGTH_EXTRA
		}
	}
	if mainLength != 0 {
		buff = make([]byte, mainLength)
		n,err := socket.Recv(buff)
		if err != nil {
			return nil,err
		}
		if n == mainLength {
			msg.SetMain(mainLength, SBuffer.Wrap(buff))
		}else{
			return nil,ERR_LENGTH_MAIN
		}
	}
	return msg
}

//发送封包
func (socket *SSocket) SendMessage(message *smessage.SMessage) bool{
	bSuccess := true
	bSuccess = socket.SendBuffer(message.HeadLength(), message.Head())
	bSuccess = socket.SendBuffer(message.ExtraLength(), message.Extra())
	bSuccess = socket.SendBuffer(message.MainLength(), message.Main())
	return bSuccess
}

func (socket *SSocket) SendBuffer(sLen int, buffer SBuffer.SBuffer) bool {
	if buffer.Limit() > 0 {
		length := 0
		for {
			n, err := socket.Write(buffer.Bytes()[length:])
			if n > 0 {
				length += n
			}
			if err != nil {
				//如果写出数据错误返回false
				break
			}
			if length == sLen {
				buffer.SetLimit(0)
				return true
			}
		}
	}
	return false
}
