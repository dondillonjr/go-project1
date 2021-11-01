package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	//Panics happen after deferred statements are executed
	panic("something bad happened")
	fmt.Println("end")
}