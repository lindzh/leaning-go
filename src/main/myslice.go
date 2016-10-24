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

	fmt.Println(" ")

	sl2 = append(sl2, 99, 98, 87)

	fmt.Println(sl2)

	sl2 = append(sl2, mySlice[4:]...)

	fmt.Println(sl2)

	fmt.Println(cap(sl2))

	fmt.Println(len(sl2))

	sil1 := []int{1, 2, 3, 4, 5, 6}

	sil2 := []int{11, 12, 13}

	copy(sil1, sil2)

	fmt.Println(sil1)

	copy(sil2, sil1)

	fmt.Println(sil2)

}
