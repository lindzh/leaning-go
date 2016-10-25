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

var once sync.Once

/**
 *
 * once 所调用的方法不能传参数
 */
func onceSetup() {
	fmt.Println("once Setup")
}

func callOnce() {
	once.Do(onceSetup)
}

func readWriteTest(rwLock *sync.RWMutex) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Println("read lokced ")
	time.Sleep(time.Duration(5) * time.Second)

}

func lockTest(rwLock *sync.RWMutex) {
	rwLock.Lock()
	defer rwLock.Unlock()
	fmt.Println("write locked")
	time.Sleep(time.Duration(1) * time.Second)
}

func readmore() {
	lock := new(sync.RWMutex)
	for i := 0; i < 20; i++ {
		go readWriteTest(lock)
	}
	for j := 0; j < 5; j++ {
		go lockTest(lock)
	}
	fmt.Println("read write lock test started------")
	time.Sleep(time.Duration(20) * time.Second)

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

	//call once test

	fmt.Println("------once-------")
	for i := 0; i < 10; i++ {
		callOnce()
	}
	fmt.Println("once call finished")

	//-----------read write lock test
	fmt.Println("read write lock-------------------------")
	readmore()

}
