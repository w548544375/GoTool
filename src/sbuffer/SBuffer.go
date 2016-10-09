package SBuffer

import "errors"

type SBuffer struct {
	buf      []byte
	position int //当前指向的位置
	cap      int //buffer的容量
}

func NewSBuffer() *SBuffer {
	buff := new(SBuffer)
	buff.cap = len(buff.buf)
	return buff
}

const (
	DEFAULT_BUFF_LENGTH = 64
)

var ErrTooLarge = errors.New("SBuffer too large")

//扩展buf
func (buff *SBuffer) expand(length int) {
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	newCap := buff.cap + length
	newBuf := make([]byte, newCap)
	newPos := copy(newBuf[:], buff.buf[:buff.position])
	buff.cap = newCap
	buff.buf = newBuf
	buff.position = newPos
}

func (buff *SBuffer) PutByte(value byte) {
	if buff.position >= buff.cap || buff.position + 1 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position] = value
	buff.position += 1
}

func (buff *SBuffer) PutByteTo(pos int, value byte) {
	if  pos+1 > buff.cap {
		//buff.expand(DEFAULT_BUFF_LENGTH)
		buff.expand(1)
	}
	buff.buf[pos] = value
}

func (buff *SBuffer) PutShort(value int16) {
	if buff.position == buff.cap || buff.position + 2 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position] = byte(value)
	buff.buf[buff.position + 1] = byte(value >> 8 & 0xFF)
	buff.position += 2
}

func (buff *SBuffer) PutShortTo(pos int, value int16) {
	if pos + 2 > buff.cap {
		buff.expand(2)
	}
	buff.buf[pos] = byte(value & 0xFF)
	buff.buf[pos + 1] = byte(value >> 8 & 0xFF)
}

func (buff *SBuffer) PutInt(value int32) {
	if buff.position == buff.cap || buff.position + 4 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position + 0] = byte(value)
	buff.buf[buff.position + 1] = byte(value >> 8 & 0xFF)
	buff.buf[buff.position + 2] = byte(value >> 16 & 0xFF)
	buff.buf[buff.position + 3] = byte(value >> 24 & 0xFF)
	buff.position += 4
}

func (buff *SBuffer) PutIntTo(pos int, value int32) {
	if pos + 4 > buff.cap {
		buff.expand(4)
	}
	buff.buf[pos] = byte(value & 0xFF)
	buff.buf[pos + 1] = byte(value >> 8 & 0xFF)
	buff.buf[pos + 2] = byte(value >> 16 & 0xFF)
	buff.buf[pos + 3] = byte(value >> 24 & 0xFF)
}

func (buff *SBuffer) PutFloat(value float32) {
	temp := int32(value)
	buff.PutInt(temp)
}

func (buff *SBuffer) PutFloatTo(pos int, value float32) {
	temp := int32(value)
	buff.PutIntTo(pos, temp)
}

func (buff *SBuffer) PutLong(value int64) {
	if buff.position == buff.cap || buff.position + 8 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position] = byte(value)
	buff.buf[buff.position + 1] = byte(value >> 8 & 0xFF)
	buff.buf[buff.position + 2] = byte(value >> 16 & 0xFF)
	buff.buf[buff.position + 3] = byte(value >> 24 & 0xFF)
	buff.buf[buff.position + 4] = byte(value >> 32 & 0xFF)
	buff.buf[buff.position + 5] = byte(value >> 40 & 0xFF)
	buff.buf[buff.position + 6] = byte(value >> 48 & 0xFF)
	buff.buf[buff.position + 7] = byte(value >> 56 & 0xFF)
	buff.position += 8
}

func (buff *SBuffer) PutLongTo(pos int, value int64) {
	if buff.position + 8 > buff.cap {
		buff.expand(8)
	}
	buff.buf[pos + 0] = byte(value)
	buff.buf[pos + 1] = byte(value >> 8 & 0xFF)
	buff.buf[pos + 2] = byte(value >> 16 & 0xFF)
	buff.buf[pos + 3] = byte(value >> 24 & 0xFF)
	buff.buf[pos + 4] = byte(value >> 32 & 0xFF)
	buff.buf[pos + 5] = byte(value >> 40 & 0xFF)
	buff.buf[pos + 6] = byte(value >> 48 & 0xFF)
	buff.buf[pos + 7] = byte(value >> 56 & 0xFF)
}

func (buff *SBuffer) PutString(value string) {
	length := len(value) + 1
	if buff.position == buff.cap || buff.position + 4 + length > buff.cap {
		buff.expand(4 + length)
	}
	buff.PutInt(int32(length))   //写入字符串长度 ，字符串的长度是c字符串的长度，也就是后面有\0
	copy(buff.buf[buff.position:], value)
	buff.buf[buff.position +length] = 0
	buff.position += length
}


func (buff *SBuffer) PutStringTo(pos int, value string) {
	length := len(value) + 1
	if buff.cap <= length + pos + 4 {
		buff.expand(length + 4)
	}
	buff.PutInt(int32(length))   //写入字符串长度 ，字符串的长度是c字符串的长度，也就是后面有\0
	copy(buff.buf[pos:], value)
	buff.buf[pos +len(value)+ 1] = 0
}

//--------------------万恶的分割线--------------------------

func (buff *SBuffer) GetByte(pos int) byte {
	return 0
}

func (buff *SBuffer) GetShort(pos int) int16 {

	return 0
}

func (buff *SBuffer) GetInt(pos int) int32 {
	return 0
}

func (buff *SBuffer) GetFloat(pos int) float32 {

	return 0
}

func (buff *SBuffer) GetLong(pos int) int64 {

	return 0
}

func (buff *SBuffer) GetString(pos int) string {

	return ""
}
