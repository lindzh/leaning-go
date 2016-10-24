package main

import (
	"fmt"
)

func main() {

	a := 99000

	if a < 100 {
		fmt.Println("a <100")
	} else if a < 1000 {
		fmt.Println("a<1000")
	} else if a < 10000 {
		fmt.Println("a<10000")
	} else {
		fmt.Println("a>10000")
	}

	fmt.Println("hahaha")

}
