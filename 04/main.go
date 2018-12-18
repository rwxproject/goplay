package main

import "fmt"

func main() {
	s := sum(1, 2, 1000)
	fmt.Println(s)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
