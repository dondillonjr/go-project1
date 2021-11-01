package main

import (
	"fmt"
)

func main() {
	//defer = runs after main func is done but
	// before main function returns
    // last in is first out
	//end
	//middle
	//start
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}