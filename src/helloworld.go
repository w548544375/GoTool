package main

import (
	"bytes"
	"fmt"
	"math"
	"sbuffer"
)

func main() {

	test := SBuffer.NewSBuffer()
	test.PutByte(24)
	test.PutShort(0x0BBA)
	test.PutInt(0x140A)
	test.PutFloat(1.0)
	fmt.Printf("%X",test)
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
	return math.Sqrt(v.x*v.x + v.y*v.y)
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
