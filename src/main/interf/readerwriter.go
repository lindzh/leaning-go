package interf

import (
	"fmt"
	"main/pojo"
)

type ReaderWriter struct {
	pojo.Person
}

func (rw *ReaderWriter) Read(buf []byte) (n int, err error) {
	fmt.Println("hahaha read")
	return 0, nil
}

func (rw *ReaderWriter) Write(buf []byte) (n int, err error) {
	fmt.Println("hahah write")
	return 0, nil
}
