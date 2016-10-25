package main

import (
	"fmt"
	"main/interf"
)

func main() {
	rw := new(interf.ReaderWriter)
	var w interf.Writer = rw
	r, ok := w.(interf.Writer)
	if ok {
		fmt.Println("ok------")
	}
	fmt.Println(r)
	rw.Name = "lindezhi"
	rw.Work()

	var myint interface{} = 1
	switch v := myint.(type) {
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case int32:
		fmt.Println("int32")
	case int64:
		fmt.Println("int64")
	case uint8:
		fmt.Println("uint8")
	case uint16:
		fmt.Println("uint16")
	case uint32:
		fmt.Println("uint32")
	default:
		fmt.Println("unknown ", v)
	}
}
