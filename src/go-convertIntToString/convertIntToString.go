package main

import (
	"fmt"
	"strconv"
)

var i int = 10
func main() {

	fmt.Println("----first----")
	var (
		x int = 5
		z string = "HELLO DON"
	)
	fmt.Println("String and int\n", x , z)
	fmt.Println("Inter I=", i)

	fmt.Println("----second----")
	//Shadowing variable (i)
	var (
		i int = 8
	)
	fmt.Println("Integer i =", i)

	fmt.Println("----third----")
	var j string
	//convert integer to string (assi)
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

	fmt.Println("----fourth----")
	var a float32 = 45.5
	var k int
	//conert float to integer
	k = int(a)
	fmt.Printf("%v, %T\n", k, k)
}