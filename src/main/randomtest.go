package main

import (
	"flag"
	"fmt"
	"math/rand"
)

const RAND_SEED = 2

/*
*
*flag 是命令行配置读取工具
 */
func main() {
	source := rand.NewSource(RAND_SEED)
	rd := rand.New(source)
	fmt.Println(rd.Int31())

	fmt.Println("flags test")

	myint := flag.Int("myint", 123, "this is a int test")

	fmt.Println("flag myint ", *myint)

}
