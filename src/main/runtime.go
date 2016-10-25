package main

import (
	"fmt"
	"runtime"
	"time"
)

// 注意闭包外面的变量  直接取i会发生改变
func goRoutine() {

	for i := 0; i < 10; i++ {
		go func(j int) {
			time.Sleep(time.Duration(j+1) * time.Second)
			fmt.Println("go routine ", j)
		}(i)
	}
}

func main() {
	fmt.Println(runtime.NumCPU())
	goRoutine()

	fmt.Println(runtime.NumGoroutine())

	time.Sleep(time.Duration(5) * time.Second)

	fmt.Println(runtime.NumGoroutine())

	time.Sleep(time.Duration(10) * time.Second)

	fmt.Println(runtime.NumGoroutine())

	time.Sleep(time.Duration(15) * time.Second)

	fmt.Println(runtime.NumGoroutine())
}
