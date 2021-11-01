package main

import (
	"fmt"
	"log"
) 

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
func panicker() {
	//anonomous deferred function - is function with no Name
	//called one time at run time
	fmt.Println("about to panic")
	defer func() {
		//recover() function returns nil if app is not panicing
		// otherwise the error causing app to panic
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()  //perens make function execute
	panic("Something bad happened")
	fmt.Println("done panicking")  //Not executed
}