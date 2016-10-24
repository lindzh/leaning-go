package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) (res int, err error) {
	if a < 0 || b < 0 {
		errors.New("params can't be less than 0")
		return
	} else {
		return a + b, nil
	}
}

func addm(a ...int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}

func funcTest() {
	ff := func(a int) (res int, err error) {
		if a < 0 {
			errors.New("invalid param a")
			return
		} else if a == 0 {
			return 0, nil
		} else {
			var result int = 1
			for i := 1; i <= a; i++ {
				result *= i
			}
			return result, nil
		}
	}

	fmt.Println(ff(5))
}

func main() {

	fmt.Println(add(10, 20))

	res, err := add(10, -1)

	if err != nil {
		fmt.Println("invalid param")
	} else {
		fmt.Println("result is ", res)
	}

	fmt.Println(addm(1, 4, 7, 9, 10))

	fmt.Println(addm([]int{1, 4, 6, 7, 8}...))

	funcTest()
}
