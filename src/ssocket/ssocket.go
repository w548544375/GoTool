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

type SSocket struct {
	net.TCPConn
}

func (socket *SSocket) Recv(buff []byte) error {
	defer socket.Close()
	length := 0
	for {
		n, err := socket.Read(buff[length:])
		if n > 0 {
			length += n
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}

//接受数据 组装为message
func (socket *SSocket) RecvMessage() *smessage.SMessage {
	msg := new(smessage.SMessage)
	buff := make([]byte, DEFAULT_HEAD_LENGTH)
	socket.Recv(buff)
	//验证数据包
	headBuf := SBuffer.Wrap(buff)
	msgType := headBuf.GetShortFrom(0)
	extraLength := headBuf.GetShortFrom(2)
	mainLength := headBuf.GetIntFrom(4)
	validate := headBuf.GetShortFrom(16)
	if msgType^extraLength^mainLength != validate {
		defer socket.Close()
		return nil
	}
	msg.SetHead(18, headBuf)
	if extraLength != 0 {
		buff = make([]byte, extraLength)
		err := socket.Recv(buff)
		if err != nil {
			defer socket.Close()
			return nil
		}
		msg.SetExtra(extraLength, SBuffer.Wrap(buff))
	}
	if mainLength != 0 {
		buff = make([]byte, mainLength)
		err := socket.Recv(buff)
		if err != nil {
			defer socket.Close()
			return nil
		}
		msg.SetMain(mainLength, SBuffer.Wrap(buff))
	}
	return msg
}

//发送封包
func (socket *SSocket) SendMessage(message *smessage.SMessage) {
	socket.SendBuffer(message.HeadLength(), message.Head())
	socket.SendBuffer(message.ExtraLength(), message.Extra())
	socket.SendBuffer(message.MainLength(), message.Main())
}

func (socket *SSocket) SendBuffer(sLen int, buffer SBuffer.SBuffer) {
	defer socket.Close()

	if buffer.Limit() > 0 {
		length := 0
		for {
			n, err := socket.Write(buffer.Bytes()[length:])
			if n > 0 {
				length += n
			}
			if err != nil {
				panic(err)
			}
			if length == sLen {
				buffer.SetLimit(0)
				break
			}
		}
	}
}
