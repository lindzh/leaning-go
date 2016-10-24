package main

import "fmt"
import "os"
import "io"
import "crypto/md5"

func main() {
	fmt.Println("hello this is a go test")
	fmt.Println("hahah")
	fmt.Println(os.Getpid())
	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("%x", h.Sum(nil))

}
