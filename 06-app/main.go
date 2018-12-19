package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hi there")
	io.WriteString(os.Stdout, "hi io wr")
}
