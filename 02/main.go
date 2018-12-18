package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type superPerson struct {
	person
	license bool
}

func main() {

	p1 := person{
		first: "abc",
		last:  "def",
		age:   23,
	}

	p2 := superPerson{
		person: person{
			first: "qwe",
			last:  "yuo",
			age:   56,
		},
		license: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//for i, v:= range x {
	//	fmt.Println(i,v)
	//}

	//fmt.Println(x[1:(len(x)-2)])

	//x = append(x,6,7,8,9)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//y := []int {11,12,13,14,15}
	//x = append(x,y...)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//
	//x = append(x[:2], x[4:]...)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//
	//
	//z := make([]int,10,100)
	//fmt.Println(z)
	//fmt.Println(len(z))
	//fmt.Println(cap(z))

	//maps

	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)
	fmt.Println(m["a"])
	//v, ok := m["c"]
	//fmt.Println(v,ok)

	if v, ok := m["a"]; ok {
		fmt.Println("key/value", v, ok)
	}

	delete(m, "a")
	fmt.Println(m)

}
