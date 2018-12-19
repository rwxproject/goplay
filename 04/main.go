package main

import "fmt"

func main() {
	s := sum(1, 2, 1000)
	fmt.Println(s)
	fmt.Println(&s)

	v := 67
	foo(&v)

}

func foo(y *int) { // pointer to an int, * means a pointer to..
	fmt.Println(y)
	fmt.Println(&y)
	*y = 45 // value to the 'y' address
	fmt.Println(y)
	fmt.Println(&y)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
