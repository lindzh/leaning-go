package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter = 0

func add(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("count:", counter, "  ", time.Now())
	lock.Unlock()
}

func main() {

	lock := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go add(lock)
	}

	for {
		lock.Lock()
		c := counter

		lock.Unlock()
		runtime.Gosched()
		if c > 8 {
			fmt.Println("hahah---:out")
			break
		}
	}
	fmt.Println("finished")

}
