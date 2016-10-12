package SBuffer

import (
	"errors"
	"unsafe"
)

type SBuffer struct {
	buf      []byte
	position int //当前指向的位置
	limit    int //位置
	cap      int //buffer的容量
}

func NewSBuffer() *SBuffer {
	buff := new(SBuffer)
	buff.cap = len(buff.buf)
	buff.limit = buff.cap
	return buff
}

const (
	DEFAULT_BUFF_LENGTH = 64
)

var ErrTooLarge = errors.New("SBuffer too large")

var ErrPosOverFlow = errors.New("设置的位置大于buff的容量")

//将byte数组转换为SBuffer对象
func Wrap(bytes []byte) *SBuffer {
	return &SBuffer{bytes, 0, len(bytes), len(bytes)}
}

func (buff *SBuffer) SetPos(newPos int) error {
	if newPos >= buff.cap {
		return ErrPosOverFlow
	}
	buff.position = newPos
	return nil
}

func (buff *SBuffer) SetLimit(newLimit int) {
	buff.limit = newLimit
}

func (buff *SBuffer) Limit() int {
	return buff.limit
}

//buff剩余没有操作的长度
func (buff *SBuffer) Len() int {
	return buff.cap - buff.position
}

//
func (buff *SBuffer) Cap() int {
	return buff.cap
}

//取得当前位置
func (buff *SBuffer) Pos() int {
	return buff.position
}

//取得到limit的buff内容
func (buff *SBuffer) Bytes() []byte {
	return buff.buf[:buff.limit]
}

//为发送做准备
func (buff *SBuffer) Flip() {
	buff.limit = buff.position
	buff.position = 0
}

//重置buff
func (buff *SBuffer) Reset() {
	buff.position = 0
	buff.buf = make([]byte, 64)
	buff.cap = len(buff.buf)
	buff.limit = buff.cap
}

/*
*每次扩展都会设置buffer的limit为cap
 */
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
	buff.limit = newCap
	buff.buf = newBuf
	buff.position = newPos
}

func (buff *SBuffer) PutByte(value byte) {
	if buff.position >= buff.cap || buff.position+1 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position] = value
	buff.position += 1
}

func (buff *SBuffer) PutByteTo(pos int, value byte) {
	if pos+1 > buff.cap {
		//buff.expand(DEFAULT_BUFF_LENGTH)
		buff.expand(1)
	}
	buff.buf[pos] = value
}

func (buff *SBuffer) PutShort(value int16) {
	if buff.position == buff.cap || buff.position+2 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position] = byte(value)
	buff.buf[buff.position+1] = byte(value >> 8 & 0xFF)
	buff.position += 2
}

func (buff *SBuffer) PutShortTo(pos int, value int16) {
	if pos+2 > buff.cap {
		buff.expand(2)
	}
	buff.buf[pos] = byte(value & 0xFF)
	buff.buf[pos+1] = byte(value >> 8 & 0xFF)
}

func (buff *SBuffer) PutInt(value int32) {
	if buff.position == buff.cap || buff.position+4 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position+0] = byte(value)
	buff.buf[buff.position+1] = byte(value >> 8 & 0xFF)
	buff.buf[buff.position+2] = byte(value >> 16 & 0xFF)
	buff.buf[buff.position+3] = byte(value >> 24 & 0xFF)
	buff.position += 4
}

func (buff *SBuffer) PutIntTo(pos int, value int32) {
	if pos+4 > buff.cap {
		buff.expand(4)
	}
	buff.buf[pos] = byte(value & 0xFF)
	buff.buf[pos+1] = byte(value >> 8 & 0xFF)
	buff.buf[pos+2] = byte(value >> 16 & 0xFF)
	buff.buf[pos+3] = byte(value >> 24 & 0xFF)
}

func (buff *SBuffer) PutFloat(value float32) {
	if buff.position == buff.cap || buff.position+4 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	bits := *(*uint32)(unsafe.Pointer(&value))
	for i := 0; i < 4; i++ {
		buff.buf[buff.position+i] = byte(bits >> uint32(i*8))
	}
	buff.position += 4
}

func (buff *SBuffer) PutFloatTo(pos int, value float32) {
	if buff.position+4 > buff.cap {
		buff.expand(4)
	}
	bits := *(*uint32)(unsafe.Pointer(&value))
	for i := 0; i < 4; i++ {
		buff.buf[pos+i] = byte(bits >> uint32(i*8))
	}
}

func (buff *SBuffer) PutLong(value int64) {
	if buff.position == buff.cap || buff.position+8 > buff.cap {
		buff.expand(DEFAULT_BUFF_LENGTH)
	}
	buff.buf[buff.position] = byte(value)
	buff.buf[buff.position+1] = byte(value >> 8 & 0xFF)
	buff.buf[buff.position+2] = byte(value >> 16 & 0xFF)
	buff.buf[buff.position+3] = byte(value >> 24 & 0xFF)
	buff.buf[buff.position+4] = byte(value >> 32 & 0xFF)
	buff.buf[buff.position+5] = byte(value >> 40 & 0xFF)
	buff.buf[buff.position+6] = byte(value >> 48 & 0xFF)
	buff.buf[buff.position+7] = byte(value >> 56 & 0xFF)
	buff.position += 8
}

func (buff *SBuffer) PutLongTo(pos int, value int64) {
	if buff.position+8 > buff.cap {
		buff.expand(8)
	}
	buff.buf[pos+0] = byte(value)
	buff.buf[pos+1] = byte(value >> 8 & 0xFF)
	buff.buf[pos+2] = byte(value >> 16 & 0xFF)
	buff.buf[pos+3] = byte(value >> 24 & 0xFF)
	buff.buf[pos+4] = byte(value >> 32 & 0xFF)
	buff.buf[pos+5] = byte(value >> 40 & 0xFF)
	buff.buf[pos+6] = byte(value >> 48 & 0xFF)
	buff.buf[pos+7] = byte(value >> 56 & 0xFF)
}

func (buff *SBuffer) PutString(value string) {
	length := len(value) + 1
	if buff.position == buff.cap || buff.position+4+length > buff.cap {
		buff.expand(4 + length)
	}
	buff.PutInt(int32(length)) //写入字符串长度 ，字符串的长度是c字符串的长度，也就是后面有\0
	copy(buff.buf[buff.position:], value)
	buff.buf[buff.position+length] = 0
	buff.position += length
}

func (buff *SBuffer) PutStringTo(pos int, value string) {
	length := len(value) + 1
	if buff.cap <= length+pos+4 {
		buff.expand(length + 4)
	}
	buff.PutInt(int32(length)) //写入字符串长度 ，字符串的长度是c字符串的长度，也就是后面有\0
	copy(buff.buf[pos:], value)
	buff.buf[pos+len(value)+1] = 0
}

//--------------------万恶的分割线--------------------------

func (buff *SBuffer) GetByte() byte {
	value := buff.buf[buff.position]
	buff.position += 1
	return value
}

func (buff *SBuffer) GetByteFrom(pos int) byte {
	return buff.buf[pos]
}

//GetShort
func (buff *SBuffer) GetShort() int16 {
	var value int16
	for i := 0; i < 2; i++ {
		value |= int16((buff.buf[buff.position+i] & 0x00FF)) << uint32(i*8)
	}
	buff.position += 2
	return value
}

func (buff *SBuffer) GetShortFrom(pos int) int16 {
	var value int16
	for i := 0; i < 2; i++ {
		value |= int16(buff.buf[pos+i]&0x00FF) << uint32(i*8)
	}
	return value
}

//GetInt
func (buff *SBuffer) GetInt() int32 {
	var value int32
	for i := 0; i < 4; i++ {
		value |= int32(buff.buf[buff.position+i]&0x000000FF) << uint32(i*8)
	}
	buff.position += 4
	return value
}

func (buff *SBuffer) GetIntFrom(pos int) int32 {
	var value int32
	for i := 0; i < 4; i++ {
		value |= int32(buff.buf[pos+i]&0x000000FF) << uint32(i*8)
	}
	return value
}

//GetFloat
func (buff *SBuffer) GetFloat() float32 {
	intValue := buff.GetInt()
	buff.position += 4
	return *(*float32)(unsafe.Pointer(&intValue))
	//bits := binary.LittleEndian.Uint32(bytes)
	//return math.Float32frombits(bits)
}

func (buff *SBuffer) GetFloatFrom(pos int) float32 {
	intValue := buff.GetIntFrom(pos)
	return *(*float32)(unsafe.Pointer(&intValue))
}

//GetLong
func (buff *SBuffer) GetLong() int64 {
	var value int64
	for i := 0; i < 8; i++ {
		value |= int64(buff.buf[buff.position+i]&0x00000000000000FF) << uint32(i*8)
	}
	buff.position += 8
	return value
}

func (buff *SBuffer) GetLongFrom(pos int) int64 {
	var value int64
	for i := 0; i < 8; i++ {
		value |= int64(buff.buf[pos+i]&0x00000000000000FF) << uint32(i*8)
	}
	return value
}

//GetString
func (buff *SBuffer) GetString() string {
	length := buff.GetInt()
	buff.position += int(length) + 4
	return string(buff.buf[buff.position : buff.position+int(length)-1])
}

func (buff *SBuffer) GetStringFrom(pos int) string {
	length := buff.GetIntFrom(pos)
	return string(buff.buf[pos+4 : pos+int(length)+4-1])
}
