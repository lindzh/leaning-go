package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int {
	fmt.Println("hahahah add ", time.Now())
	return a + b
}

func main() {
	go add(1, 2)
	fmt.Println("dddddd", time.Now())
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("hahahha ----- ", time.Now())
}
