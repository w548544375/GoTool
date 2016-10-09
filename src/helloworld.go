package main

import (
	"fmt"
	"math"
	"bytes"
	"sbuffer"
)

func main() {

	test := SBuffer.NewSBuffer()
	test.PutByte(24)
	test.PutShort(0x0BBA)
	test.PutInt(0x140A)
	test.PutFloat(1.0)
	test.PutLong(85776868)
	test.PutString("佛啊佛大佛杀佛见哦企鹅技巧哦见哦天辅导费激动啊就佛大家哦大家供电局阿公大家感动啊更激动啊激动感觉哦啊发激动啊发激动啊附近的")
	buff := bytes.NewBufferString("wang")
	fmt.Println(buff.Bytes())
	fmt.Printf("%v", test)
}

type MyBuff struct {
	bytes.Buffer
}

type Vertex struct {
	x, y float64
}

func (v  *Vertex ) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("unsupported number.%v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}