package main

import "fmt"

// func (r receiver) identifier(parameters) (return(s)) { code... }

// keyword indentifier type
// a value can be more than one type

type foo struct {
	fn string
}

type human interface {
	speak()
}

func bar(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	for i, v := range x {
		fmt.Println(i, v)
	}

}

func (f foo) speak(s string) {
	fmt.Println("foo said ", s)
}

func main() {
	bar(1, 2, 3, 4, 5)
	ab := foo{
		fn: "afn",
	}
	ab.speak("whaaat")
}
