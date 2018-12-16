package main

import (
	"fmt"
	"runtime"
)

var x int
var y string
var z bool

type hd int
var a hd

func main() {

	x = 43
	y = "Jamie"
	z = true

	fmt.Println(x,y,z)
	s := fmt.Sprintf("%v %v %v",x,y,z)
	fmt.Println(s)

	fmt.Printf("%T\n",a)
	a = 42
	fmt.Println(a)

	b := int(a)
	fmt.Printf("%T\n",b)
	fmt.Println(b)

	fmt.Println(runtime.GOOS)
	fmt.Printf(runtime.GOARCH)

	w := `ABC`
	fmt.Printf("%T\n",w)
	bs := []byte(w)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%#U ",s[i])
	//}

	l := 5457
	fmt.Printf("%d\t%b\t%#x\n",l,l,l)

	m := l << 1
	fmt.Printf("%d\t%b\n",m,m)

	n := m >> 1
	fmt.Printf("%d\t%b\n",n,n)


}
