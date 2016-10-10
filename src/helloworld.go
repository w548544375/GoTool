package main

import (
	"bytes"
	"fmt"
	"math"
	"sbuffer"
)

func main() {

	test := SBuffer.NewSBuffer()
	test.PutShort(0x0BBA)
	sd := test.GetShortFrom(0)
	fmt.Printf("取出的Short为：%04X\n", sd)
	test.PutInt(1)
	intval := test.GetIntFrom(2)
	fmt.Printf("取出的Int为：%d\n", intval)
	test.PutFloat(1.0)
	floatValue := test.GetFloatFrom(6)
	fmt.Printf("取出的Float为：%f\n", floatValue)
	test.PutString("HelloWorld")
	st := test.GetStringFrom(10)
	fmt.Printf("取出的String为：%v\n", st)
	fmt.Printf("%v", test)
}

type MyBuff struct {
	bytes.Buffer
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) {
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
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
