package ssession

import (
	"errors"
	"io"
	"net"
	"sbuffer"
	"smessage"
)

const (
	DEFAULT_HEAD_LENGTH = 18
)

var LENGHT_ERROR = errors.New("Packet Length mismatch")
var ERR_LENGTH_EXTRA = errors.New("Extra Packet Length mismatch")
var ERR_LENGTH_MAIN = errors.New("Main Packet Length mismatch")

type SSocket struct {
	conn net.Conn
}

func (self *SSocket) Close() {
	self.conn.Close()
}

func (socket *SSocket) Recv(buff []byte) (int, error) {
	length := 0
	for {
		n, err := socket.conn.Read(buff[length:])
		if n > 0 {
			length += n
		}
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return length, nil
}

//接受数据 组装为message
func (socket *SSocket) RecvMessage() (*smessage.SMessage, error) {
	msg := new(smessage.SMessage)
	buff := make([]byte, DEFAULT_HEAD_LENGTH)
	socket.Recv(buff)
	//验证数据包
	headBuf := sbuffer.Wrap(buff)
	msgType := headBuf.GetShortFrom(0)
	extraLength := headBuf.GetShortFrom(2)
	mainLength := headBuf.GetIntFrom(4)
	validate := headBuf.GetShortFrom(16)
	if int(msgType)^int(extraLength)^int(mainLength) != int(validate) {
		return nil, LENGHT_ERROR
	}
	msg.SetHead(DEFAULT_HEAD_LENGTH, headBuf)
	if extraLength != 0 {
		buff = make([]byte, extraLength)
		n, err := socket.Recv(buff)
		if err != nil {
			return nil, err
		}
		if n == int(extraLength) {
			msg.SetExtra(extraLength, sbuffer.Wrap(buff))
		} else {
			return nil, ERR_LENGTH_EXTRA
		}
	}
	if mainLength != 0 {
		buff = make([]byte, mainLength)
		n, err := socket.Recv(buff)
		if err != nil {
			return nil, err
		}
		if n == int(mainLength) {
			msg.SetMain(int(mainLength), sbuffer.Wrap(buff))
		} else {
			return nil, ERR_LENGTH_MAIN
		}
	}
	return msg, nil
}

//发送封包
func (socket *SSocket) SendMessage(message *smessage.SMessage) bool {
	bSuccess := true
	bSuccess = socket.SendBuffer(int(message.HeadLength()), message.Head())
	bSuccess = socket.SendBuffer(int(message.ExtraLength()), message.Extra())
	bSuccess = socket.SendBuffer(int(message.MainLength()), message.Main())
	return bSuccess
}

func (socket *SSocket) SendBuffer(sLen int, buffer *sbuffer.SBuffer) bool {
	if buffer.Limit() > 0 {
		length := 0
		for {
			n, err := socket.conn.Write(buffer.Bytes()[length:])
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

func NewSocket(client net.Conn) *SSocket {
	return &SSocket{conn: client}
}
