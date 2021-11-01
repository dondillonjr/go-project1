package main

import (
	"fmt"
)

var i int = 10
func main() {

	const (
		n = iota  // is an integer who is to be incremented
		m = iota
		o = iota
	)

	const (
		a = iota
		b
		c
	)

	const (
		_ = iota  //assigns 0 to _ which is thrown away
		d
		e
	)

	const (
		_ = iota + 5  //thrown away
		cat  // has value 6
		dog  // has value 7
		fish // has value 8
	)
    fmt.Println(n)
	fmt.Println(m)
	fmt.Println(o)
	fmt.Println("watch1")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) 
	fmt.Println("watch2")
	fmt.Println(d)
	fmt.Println(e) 
	fmt.Println("watch3")
	fmt.Println(cat)
	fmt.Println(dog)
	fmt.Println(fish)
}


	

