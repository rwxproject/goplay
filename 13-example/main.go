package main

import (
	"fmt"

	"github.com/rwxproject/goplay/13-example/acdc"
)

func main() {
	fmt.Println(acdc.Sum(1, 2, 3, 4, 5))
}

// to run benchmark run:
// go test -bench .

// to test run
// go test
