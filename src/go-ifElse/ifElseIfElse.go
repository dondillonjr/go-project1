package main

import (
	"fmt"
	"math"
)

func main() {
	// Create MAP
	// MAP Key - Value
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        22343433,
		"Florida":      43454544,
		"New York":     14423778,
		"Pennsylvania": 52321234,
		"Ohio":         67668900,
	}

	fmt.Println("----first----")
       // initialize -----------------------|  test
	   // grab the value for key Flordia
	   // check if Flordia exist ~ ok = true
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop, ok)
	}
	
	fmt.Println("----second----")
	number := 50
	guess  := 30
    // Compairson Operators
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it:")
		}

		// Compairson Operators
		fmt.Println(number<=guess, number >=guess, number!=guess)
	} else if guess < 1 || guess > 100 {
		fmt.Println("The guess must be less than 1 or greater than 100")
	} else {
		fmt.Println()
	}
	
	fmt.Println("----third----")
	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(false)
	fmt.Println(!false)

	fmt.Println("----fourth----")
	if returnTrue() {
		fmt.Println("Calling returnTrue function")		
	}

	fmt.Println("----fifth----")
	mynum := 0.1 
	//get square-root of num then square it
	//Note you get the same number
	if mynum == math.Pow(math.Sqrt(mynum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	fmt.Println("----sixth----")
	myNum := 0.12356789
	//get square-root of num then square it
	// if you add decimals you must take absolute value and subtract (1)
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}