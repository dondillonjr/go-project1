package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----first------")
	switch 3 {
	case 1 , 2, 3:
		fmt.Println("one, two or three")
	case 4, 5 ,6:
		fmt.Println("four, five or six")
	default:
		fmt.Println("some other number")
	}
	fmt.Println()
	fmt.Println("-----second------")
    // 4+4=8, ;i is tag tested against
	switch i:= 4 + 4;i {
	case 1 , 2, 3:
		fmt.Println("one, two or three")
	case 4, 5 ,6:
		fmt.Println("four, five or six")
	default:
		fmt.Println("some other number")
	}
	fmt.Println()
	fmt.Println("-----third------")
	//Tagless syntax
	j := 20
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		fmt.Println("j <=10")
	case j <= 20:
		fmt.Println("less than or equal to twenty")
		fmt.Println("j <=20")
		fallthrough // will make code also to to next case
	case j <= 30:
		fmt.Println("less than or equal to thirty")
		fmt.Println("j <=30")
	default:
		fmt.Println("greater than thirty")
		fmt.Println("j > 30")
	}
}