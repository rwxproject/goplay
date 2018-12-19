package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		c <- 56
		c <- 89
	}()

	fmt.Println(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	// c := make(chan int, 1)
	// c <- 42

}
