package main

import (
	"fmt"
)

func main() {
	fmt.Println("----first----")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} 

	fmt.Println()
	fmt.Println("----second----")
	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("----third----")
	for i, j := 0, 0; i < 5 ; i, j = i + 1, j + 2 {
		fmt.Println(i, j)
	} 

	fmt.Println()
	fmt.Println("----fourth----")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i % 2 == 0 {
			i /= 2
		} else {
			i = 2 * i + 1
		}
	}

	fmt.Println()
	fmt.Println("----fifth----")
	i := 0
	for  ; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("----sixth----")
	k := 0
	for  ; k < 4; {
		fmt.Println(k)
		k++
	}

	fmt.Println()
	fmt.Println("----seventh----")
	c := 0
	for  ; c < 4; {
		fmt.Println(c)
		c++
	}

	fmt.Println()
	fmt.Println("----eigth----")
	d := 0
	for  ; c < 4; {
		fmt.Println(d)
		d++
		if d == 6 {
			break
		}
	}

	fmt.Println()
	fmt.Println("----ninth----")
	for i :=0 ; i < 10; i++ {
		if i%2  == 0 {
			continue
		}
		//show odd numbers
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println("----tenth----")
	for i := 1; i <=3; i++ {
		fmt.Println("A",i)
		for j := 1; j <=3; j++ {
			fmt.Printf("B %v %v\n", j, i * j)
			if i * j >= 3 {
				fmt.Println("GETs Out Inner Loop")
				break
			}
		}	
	}

	fmt.Println()
	fmt.Println("----eleventh----")
	//label
	Loop:
		for i := 1; i <=3; i++ {
			fmt.Println("A",i)
			for j := 1; j <=3; j++ {
				fmt.Printf("B %v %v\n", j, i * j)
				if i * j >= 3 {
					fmt.Println("GETs Out Inner and Outer Loop")
					break Loop  //break out at label
				}
			}	
		}

	fmt.Println()
	fmt.Println("----twelfth----")
	//Collections with for loops
	//Slice with 3 members
	s := []int{1,2,3}
	fmt.Println("Print Slice")
	fmt.Println(s)
	fmt.Println(s[1])
	//range yeilds key and value
	// index & value
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println()
	fmt.Println("----thirteenth----")
	//Collections with for loops
	//Array with three members
	y := [3]int{1,2,3}
	fmt.Println("Print Array")
	fmt.Println(y)
	fmt.Println(y[1])
	//range yeilds key and value
	// index & value
	for k, v := range y {
		fmt.Println(k, v)
	}

	fmt.Println()
	fmt.Println("----fourteenth----")
	//Collections with for loops
	//Map with 5 fields
	statePopulation := map[string]int{
		"Louisiana": 32233333,
		"Texas":     55443444,
		"Alabama":   33434344,
		"Mississippi": 2222323,
		"Viginia":    56699669,
	}
	fmt.Println("Print Map")
	fmt.Println(statePopulation)
	fmt.Println()
	fmt.Println(statePopulation["Texas"])
	//range yeilds key and value
	// index & value
	// Key & Value
	fmt.Println()
	fmt.Println("----fifteenth----")
	for k, v := range statePopulation {
		fmt.Println(k, v)
	}
	// Key Only
	fmt.Println()
	fmt.Println("----sixteenth----")
	for k := range statePopulation {
		fmt.Println("Key ONLY", k)
	}

	// Value Only
	fmt.Println()
	fmt.Println("----seventeenth----")
	for _, v := range statePopulation {
		fmt.Println("Value ONLY", v)
	}

	fmt.Println()
	fmt.Println("----eighteenth----")
	t := "Hello DON"
	//goes through one char at a time
	//index with unicode char value 
	for k, v := range t {
		fmt.Println(k, v)
	}

	fmt.Println()
	fmt.Println("----nineteenth----")
	z := "Hello DON"
	//goes through one char at a time
	//index with string char 
	for k, v := range z {
		fmt.Println(k, string(v))
	}
}