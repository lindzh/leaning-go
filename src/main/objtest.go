package main

import (
	"fmt"
	"main/pojo"
)

func main() {
	person := new(pojo.Person)
	person.Name = "lindezhi"
	person.Id = 12345
	person.Age = 10
	person.Work()
	fmt.Println("invoke obj and package success")
	person.SayHello()
	fmt.Println(person.Age)

	// object inner zuhe test
	var worker = new(pojo.Worker)
	worker.Name = "lindezhi"
	worker.Id = 1000
	worker.Age = 21
	worker.Job = "coding"
	fmt.Println(worker)
	worker.Work()
	fmt.Println(worker.Job)
	worker.SayHello()
}
