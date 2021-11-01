package main

import (
	"fmt"
)

func main() {
	fmt.Println("---one---")
	d := divide(5.0, 3.0)
	fmt.Println(d)

	fmt.Println("---two---")
	f, err := divide2( 9.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

	fmt.Println("---three---")
	c, err := divide2( 5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
}

func divide(a, b float64) float64 {
	return a / b
}

func divide2(a, b float64) (float64, error)  {
	if b == 0.0 {
		//panic("Cannot provide zero as second value")
		return 0.0, fmt.Errorf("Cannot devide by zero")
	}
	return a / b, nil
}