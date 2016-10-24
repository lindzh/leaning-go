package main

import (
	"fmt"
)

func main() {

	a := 10

	switch a {
	case 8:
		fmt.Println("a=8")
	case 9:
		fmt.Println("a=9")
	case 10:
		fmt.Println("a=10")
	default:
		fmt.Println("hahah ignoreed")
	}

	var ss string = "a"

	switch ss {
	case "b":
		fmt.Println("ss=b")
	case "c":
		fmt.Println("ss=c")
	case "a":
		fmt.Println("ss=a")
	}

}
