package main

import (
	"fmt"
)

func main() {

	ss := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ss)

	for i := 0; i < len(ss); i++ {
		fmt.Print(ss[i], " ")
	}

	fmt.Println("")

	for i, v := range ss {
		fmt.Println(i, "  ", v)
	}

	testmap := map[string]string{
		"name": "lin",
		"age":  "feee",
		"haha": "test",
	}

	for k, v := range testmap {
		fmt.Println(k, "  ", v, "  ", testmap[k])
	}

}
