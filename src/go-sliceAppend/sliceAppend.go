package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----first------")
	j := []int{1,2,3,4,5}
	//Pop off elements in middle of slice
	//pop off element after index 1
	//pop off element at index 4 and before
	fmt.Println("\npops off elements in middle of slice")
	//j[:2]= 1,2
	//j[5:]= 6,7,8,9,10
	fmt.Printf("J=%v\n", j)
	fmt.Println("-----second------")
	p := append(j[:2], j[3:]...)
	fmt.Printf("J=%v\n", j)
	fmt.Printf("P=%v\n", p)	
}