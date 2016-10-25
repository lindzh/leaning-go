package main

import (
	"errors"
	"fmt"
	"time"
)

/**
 * chan超时使用另外一个chan读取数据，使用go routine启动另一个sleep后释放
 *
 */
func timeout(sec int) (ret chan bool, err error) {
	if sec < 1 {
		errors.New("invalid time")
		return
	}
	t := make(chan bool)
	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		t <- true
	}()
	return t, nil
}

func wait(ch chan int, tt int) (ret int, err error) {
	tch, e := timeout(tt)
	if e != nil {
		errors.New("init timeout failed")
		return
	} else {
		select {
		case <-tch:
			fmt.Println("time out")
			return 0, nil
		case res := <-ch:
			fmt.Println("readed ok")
			return res, nil
		}
	}
}

type MyChan struct {
	ich <-chan int
	och chan<- int
}

/*
 *
 * read chan 需要用括号
 */
func newIntChan() MyChan {
	var ch chan int = make(chan int)

	var ich (<-chan int) = (<-chan int)(ch)
	var och chan<- int = chan<- int(ch)
	return MyChan{ich, och}
}

func iochanTest() {
	var ch MyChan = newIntChan()
	go func(out chan<- int) {
		time.Sleep(time.Duration(3) * time.Second)
		fmt.Println("write chan 10")
		out <- 10
		fmt.Println("write chan 20")
		out <- 20
	}(ch.och)

	go func(in <-chan int) {
		ret1 := <-in
		fmt.Println("read chan :", ret1)
		ret2 := <-in
		fmt.Println("read chan :", ret2)
	}(ch.ich)

	fmt.Println("chan in out started")
	time.Sleep(time.Duration(10) * time.Second)

}

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		ch <- 20
	}()
	ret, err := wait(ch, 3)
	if err != nil {
		fmt.Println("chan wait error")
	} else {
		fmt.Println("chan wait ok ", ret)
	}

	//-----------------------------
	fmt.Println("io chan start --------------")
	iochanTest()

}
