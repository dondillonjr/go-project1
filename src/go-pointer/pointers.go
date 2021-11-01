package main

import (
	"fmt"
)

//Pointers
func main() {
	fmt.Println("-----first-----")
	a := 42
	b := a
	fmt.Println(a , b)
	fmt.Println("-----second-----")
	a = 27
	fmt.Println(a, b)

	fmt.Println()
	fmt.Println("-----third-----")
	var c int = 20
	//d is a pointer to integer
	var d *int = &c
	fmt.Println(c ,d)
	fmt.Println(&c ,d)
	fmt.Println(c ,*d) //*d de-reference
	fmt.Println()
	fmt.Println("-----fourth-----")
	c = 19
	fmt.Println(c, *d)
	fmt.Println("-----fifth-----")
	*d = 3
	fmt.Println(c, *d)
}