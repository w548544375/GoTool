package main

import (
	"fmt"
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


