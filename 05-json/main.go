package main

// https://mholt.github.io/json-to-go/

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "fname",
		Last:  "lname",
		Age:   32,
	}
	p2 := person{
		First: "faname",
		Last:  "laname",
		Age:   33,
	}
	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	var speople []person

	ss := `[{"First":"fname","Last":"lname","Age":32},{"First":"faname","Last":"laname","Age":33}]`
	bss := []byte(ss)
	fmt.Println(bss)

	err1 := json.Unmarshal(bss, &speople)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(speople)
}
