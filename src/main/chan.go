package main

import (
	"fmt"
	"time"
)

func add(ch chan int, index int) {
	//write wait
	ch <- 1
	fmt.Println("hahah--- ", index)
}

func readChanBuf(c chan int) {
	var cc int = 0
	for v := range c {
		fmt.Println("chan :", v)
		cc++
		if cc > 8 {
			break
		}
	}
	fmt.Println("read chan buf ------")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add(chs[i], i)
	}

	for _, c := range chs {
		//read wait
		<-c
	}

	fmt.Println("chan finished")
	time.Sleep(time.Duration(2) * time.Second)

	//=============== chan buf test
	chb := make(chan int, 20)

	for j := 0; j < 19; j++ {
		chb <- j
	}
	readChanBuf(chb)

	fmt.Println("read chb finished")

	go func(ch chan int) {
		for id := 20; id < 40; id++ {
			ch <- id
		}
	}(chb)

	fmt.Println("add more now")
	var cc int = 0

	for {
		cv := <-chb
		fmt.Println("chb: ", cv)
		cc++
		if cc >= 30 {
			break
		}
	}
	fmt.Println("chan buf test finish")

}
