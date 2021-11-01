package main

import (
	"fmt"
)

//Pointers
func main() {
	a := [3]int{1,2,3}
	b := &a[0]
	c := &a[1]
	// %p is value of pointer
	fmt.Printf("%v %p %p\n", a, b, c)
}