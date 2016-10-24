package main

import (
	"fmt"
)

func modify(arr [6]int) {
	arr[0] = 100
	fmt.Println(arr)
}

func modifySlice(arr []int) {
	arr[0] = 100
	fmt.Println(arr)
}

func main() {
	var a = 10

	fmt.Println(a)

	var inta int = 100

	fmt.Println(inta)

	b := 300

	inta += b

	fmt.Println(inta)

	ss := "this is a test case"

	fmt.Println(ss)

	ss2 := 'a'

	fmt.Println(ss2)

	var a8 int8 = 123

	fmt.Println(a8)

	arr1 := [6]int{1, 2, 3, 5, 6, 7}

	fmt.Println(arr1)

	fmt.Println(len(arr1))

	for idx, value := range arr1 {
		fmt.Println("idx:", idx, " value:", value)
	}

	fmt.Println(arr1[2])

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])

	}

	modify(arr1)

	fmt.Println(arr1)

	mySlice := arr1[:]

	fmt.Println(mySlice)

	modifySlice(mySlice)

	fmt.Println(mySlice)

	sl2 := make([]int, 20)

	fmt.Println(sl2)

	for j := 0; j < len(sl2); j++ {

		sl2[j] = 100*j + 122
	}

	fmt.Println(sl2)

	for _, v := range sl2 {
		fmt.Print(v)
		fmt.Print(" ")
	}
}
