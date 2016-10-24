package main

import (
	"fmt"
)

/**
 *
 * 现在竟然可以了
 */
func ifelseTest(a int) bool {
	if a > 100 {
		return true
	} else {
		return false
	}
}

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

	fmt.Println(ifelseTest(100))

}
