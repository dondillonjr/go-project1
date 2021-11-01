package main

import (
	"fmt"
)

 type myStruct struct {
	 foo int
 }

 type myStruct2 struct {
	foo int
}

func main() {
	fmt.Println("-----first------")
	//ms is a var of type myStruct
	var ms myStruct
	ms = myStruct{foo: 42}
	fmt.Println(ms)
	fmt.Println(ms.foo)

	fmt.Println("-----second------")
	var ss *myStruct2
	//ss is holding an address of an object that has a field with a value of 22
	ss = &myStruct2{foo: 22}
	fmt.Println(ss)
	fmt.Println(ss.foo)

	fmt.Println("-----third------")
	var kk *myStruct2
	fmt.Println(kk)
}