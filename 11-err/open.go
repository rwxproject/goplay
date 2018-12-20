package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("names1.txt")
	if err != nil {
		// fmt.Errorf("additional info ...: %v", err)
		log.Println(err)
		// log.Fatalln(err)
		// log.Panicln(err)
		// fmt.Errorf("additional info ...: %v", err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
