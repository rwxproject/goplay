package main

// to see documentation run:
//  godoc -http=:8080

import (
	"fmt"

	"github.com/rwxproject/goplay/12-pkg/abc"
)

type person struct {
	name string
	age  int
}

func main() {
	john := person{
		name: "John",
		age:  abc.Years(4),
	}
	fmt.Println(john)
}
