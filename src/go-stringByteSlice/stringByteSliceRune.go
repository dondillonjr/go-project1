package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----first------")
	s1 := "First String "
	s2 := "Second String"
	fmt.Printf("%v, %T\n", s1, s1)
	fmt.Println("-----second------")
	// yeilds bytes
	fmt.Printf("\n%v, %T\n", s1[2], s1[2])
	fmt.Println("-----third------")
	// convert byte to string
	fmt.Printf("\n%v, %T\n", string(s1[2]), s1[2])
	// can concat strings
	fmt.Println("-----fourth------")
	fmt.Printf("\n%v, %T\n", s1 + s2, s1 + s2)
	//put string in collection of bytes

	fmt.Println("-----fifth------")
	b := []byte(s1)
	fmt.Println("Store string as collection of byte slices")
	fmt.Printf("\n%v, %T\n", b, b)

	//Runes - alias for int32
	fmt.Println("-----sixth------")
	var q rune = 'a'
	r := 'a'
	fmt.Printf("\n%v, %T\n", q, q)
	fmt.Printf("%v, %T\n", r, r)
}