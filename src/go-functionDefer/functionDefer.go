package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	// defer = runs after main func is done but
	// before main function returns
	defer fmt.Println("middle")
	fmt.Println("end")
}