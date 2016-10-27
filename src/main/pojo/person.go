package pojo

import (
	"fmt"
)

type Person struct {
	Id   int64 `ffff`
	Name string
	Age  int8
}

func (person *Person) Work() {
	fmt.Println(person.Name, " is working now")
}

func (person *Person) SayHello() {
	fmt.Println("hello, I'm ", person.Name, " and I am ", person.Age, " years old")
}
