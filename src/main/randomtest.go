package main

import (
	"fmt"
	"math/rand"
)

const RAND_SEED = 2

func main() {
	source := rand.NewSource(RAND_SEED)
	rd := rand.New(source)
	fmt.Println(rd.Int31())

}
