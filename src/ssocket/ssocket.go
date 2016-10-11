package ssocket

import (
	"net"
	"smessage"
	"io"
	"sbuffer"
)

const (
	DEFAULT_HEAD_LENGTH = 18
)

type SSocket struct {
	net.TCPConn
}

func (socket *SSocket) Recv(buff []byte) {
	defer socket.Close()
	length := 0
	for {
		n, err := socket.Read(buff[length:])
		if n > 0 {
			length += n
		}
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
	}

}


//发送封包
func (socket *SSocket) SendMessage(message *SMessage.SMessage){
    socket.SendBuffer(message.HeadLength(),message.Head())
    socket.SendBuffer(message.ExtraLength(),message.Extra())
    socket.SendBuffer(message.MainLength(),message.Main())

}

func (socket *SSocket) SendBuffer(sLen int,buffer SBuffer.SBuffer){
	defer socket.Close()

	if buffer.Limit() > 0 {
		length := 0
		for {
		  n,err := socket.Write(buffer.Bytes()[length:])
			if n >0 {
				length += n
			}
			if err != nil{
				panic(err)
			}
			if length == sLen {
				buffer.SetLimit(0)
				break
			}
		}
	}
}