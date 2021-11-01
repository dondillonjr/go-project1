package main

import (
	"fmt"
)

func main() {
	// Create MAP
	// MAP Key - Value
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        22343433,
		"Flordia":      43454544,
		"New York":     14423778,
		"Pennsylvania": 52321234,
		"Ohio":         67668900,
	}
	fmt.Println("---first---")
	fmt.Println("show MAP - key & value pairs")
	//Slice cannot be a key to a map
	fmt.Println(statePopulations)
	fmt.Println("---second---")
	//map called m - map of 3 int value string pairs 
	//array is a valid key type for a map
	m := map[[3]int]string{ }
	fmt.Println(statePopulations, m)
	fmt.Println("---third---")
	// Create MAP using MAKE function - no entries
	georgiaCityPopulations := make(map[string]int)
	fmt.Println("\nMap declaration with no entries")
	fmt.Println(georgiaCityPopulations)

	fmt.Println("---fourth---")
	// Get item from MAP
	fmt.Println("\nGet \"Ohio\" item from MAP")
	fmt.Println(statePopulations["Ohio"])

	fmt.Println("---fifth---")
	// Add item to MAP
	fmt.Println("\nAdd \"Georgia\" item to MAP")
	fmt.Println(statePopulations)
	fmt.Println("---sixth---")
	statePopulations["Georgia"] = 8343434
	//Note: Order of items is not ganrentied
	fmt.Println(statePopulations)

	fmt.Println("---seventh---")
	// Delete item from MAP
	fmt.Println("\nDelete \"Georgia\" item to MAP")
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	
	fmt.Println("---eight---")
	// Check if key is in MAP
	fmt.Println("\nCheck if Key \"Georgia\" is in MAP")
	fmt.Println(statePopulations)
	_, ok := statePopulations["Georgia"]
	fmt.Println(ok)
	fmt.Println("Check if Key \"Ohio\" is in MAP")
	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)

	fmt.Println("---ninth---")
	// How many elements are in MAP
	fmt.Println("\nHow many elements are in MAP")
	fmt.Println(len(statePopulations))

	fmt.Println("---tenth---")
	// MAPs are passed by reference
	sp := statePopulations
	fmt.Println("\nMAPs are pass by references")
	fmt.Println("Note: Ohio is removed in sp and statePopulations")
	fmt.Printf("ORIG=%v\n", statePopulations)
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
} 