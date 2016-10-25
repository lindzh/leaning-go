package main

import (
	"fmt"
)

func foreach(unitl int) {
	var a int = 1
START:
	fmt.Print(a, " ")
	a++
	if a < unitl {
		goto START
	}
	fmt.Println("end -------")
}

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

	foreach(10)

	for dd, v := range ss {
		r := dd % 2
		if r == 0 {
			continue
		}
		fmt.Println(v)

	}

	fmt.Println("hahahha")

}
