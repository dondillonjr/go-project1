package main

import (
	"fmt"
)

func main() {
	
	//anonomous function
	//function declared as type
	//Note: func() is declared and invoked at same time
	fmt.Println("----first-----")
	func() {
		msg := "Hello Go"
		fmt.Println(msg)
	}()

	fmt.Println("----second-----")
	for i := 0; i < 5; i++ {
		//Inner function uses outer scope variable (i)
		//This is not best practice.
		//
		//Problems result if run asynchronously
		//
		//Asynchronous means that you can execute multiple things at 
		//a time and you don't have to finish executing the current 
		//thing in order to move on to next one.
		//
		//Best Practice is next anonomous function
		func() {
			fmt.Println(i)
		}()	
	}	

	fmt.Println("----third-----")
	for i := 0; i < 5; i++ {
		//Outer scope variaable (k) above is passed into inner 
		//scope function below.
		//
		//Best practice is to use this method and not 
		//use outer scope variable (k) shown above
		func(i int) {
			fmt.Println(i)
		}(i)	
	}	

	fmt.Println("----fourth-----")
	//function defined as variable
	f := func() {
		fmt.Println("HELLO DON")
	}	
	f()	//invoke function f

	fmt.Println("----fifth-----")
	//formally declair function as variable
	var g func() = func() {
		fmt.Println("HELLO DON")
	}	
	g()	

	fmt.Println("----sixth-----")
	//Declaring a function signiture for a Divide function
	//function takes two float input parameters
	//Returns a float and error
	//
	//NOTE: divide function is declaired as a variable
	var divide func(float64, float64) (float64, error) 
	//initialize variable (divide), set equal to anonomous function
	//that takes variable (a) and (b) which are floats and returns a
	//float and error
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("CanNot divide bu zero")
		} else {
			return a / b, nil
		}
	}	
	d, err := divide(5.0, 3.0)
	//if problem was found (xxx/0), print error and get out
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	fmt.Println("----seventh-----")
	k, err := divide(6.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(k)
}