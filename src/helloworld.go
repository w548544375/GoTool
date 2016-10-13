package main

import (
	"fmt"
)

func main() {
	// test := SBuffer.NewSBuffer()
	// test.PutShort(0x0BBA)
	// sd := test.GetShortFrom(0)
	// fmt.Printf("取出的Short为：%04X\n", sd)
	// test.PutInt(1)
	// intval := test.GetIntFrom(2)
	// fmt.Printf("取出的Int为：%d\n", intval)
	// test.PutFloat(1.0)
	// floatValue := test.GetFloatFrom(6)
	// fmt.Printf("取出的Float为：%f\n", floatValue)
	// test.PutString("HelloWorld")
	// st := test.GetStringFrom(10)
	// fmt.Printf("取出的String为：%v\n", st)
	// fmt.Printf("%v\n", test)
	// test1 := Test()
	// fmt.Printf("%v", test1)
	Test(1, 2, 3, 4)
	a := make([]int, 0)
	fmt.Printf("%v", a)
	a = append(a, 1)
	fmt.Printf("%v %d", a, len(a))

}

func Test(a ...int) {
	for i := range a {
		fmt.Printf("%d", i)
	}
	// for j := range c {
	// 	fmt.Printf("%s", j)
	// }
}
