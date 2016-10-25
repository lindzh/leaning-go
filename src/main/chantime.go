package main

import (
	"errors"
	"fmt"
	"time"
)

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

}
