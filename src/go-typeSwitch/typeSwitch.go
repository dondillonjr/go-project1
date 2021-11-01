package main

import (
	"fmt"
)

func main() {
	//type switch
	// interface can be anything
	// int, string, boolean, function, struct, collection
	var i interface{} = 1 
	j := 4
	//var i interface{} = 1.0
	//var i interface{} = "1"
	//var i interface{} = [3]int{} 
	//var i interface{} = [2]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		if j == 4 {
			fmt.Printf("J=%v\n", j)
			break
		}
		fmt.Println("String after break")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a stromg")
	case [3]int:
		fmt.Println("i is [3]int")
	case [2]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another type")
	}
}