package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	fmt.Println("-----first------")
    var ms *myStruct
	fmt.Println(ms)	 
	fmt.Println("-----second------")
	//initialize an empty myStruct object
	ms = new(myStruct)
	//ms is holding an address of an object that has a field with no value
	fmt.Println(ms)
	fmt.Println(ms.foo)

	fmt.Println("-----third------")
	ms.foo = 7
	fmt.Println(ms.foo)
}