package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----first------")
   //variantic parameters - are a list pa
   //passed in an put them in a Slice.
   //Can only have One list
	sum(1, 2, 3, 4, 5)

	fmt.Println("-----second------")
	//Variatic Parameter must be last and can only have one
	sum2("The sum is: ", 1, 2, 3, 4, 5)

	fmt.Println("-----third------")
	s := sum3(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", s)

	fmt.Println("-----fourth------")
	p := sum4(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *p)

	fmt.Println("-----fifth------")
	q := sum5(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", q)
}

//variantic parameters
//take arguments passed in an put them in a Slice
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

func sum2(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

func sum3(values ...int)  int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sum4(values ...int)  *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sum5(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}