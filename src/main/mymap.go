package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  uint8
	Work string
}

func main() {
	personDB := make(map[string]Person)

	personDB["lin"] = Person{"lin", 10, "Student"}

	fmt.Println(personDB)

	var p Person = Person{}

	p.Name = "hahah"

	p.Work = "5545"

	fmt.Println(p)

	r, find := personDB["lin"]

	if find {
		fmt.Println("find result:", r.Name)
	} else {
		fmt.Println("not found")
	}

	var mymap map[int32]string

	mymap = map[int32]string{
		32: "32323",
		43: "hahaha",
	}

	fmt.Println(mymap)

	mymap[44] = "54545"

	fmt.Println(mymap)

	fmt.Println(len(mymap))

	for k, v := range mymap {
		fmt.Println(k, v)
	}

	delete(mymap, 43)

	fmt.Println(mymap)

	delete(mymap, 1) //panic

	fmt.Println(mymap)

}
