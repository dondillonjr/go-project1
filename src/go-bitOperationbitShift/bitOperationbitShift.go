package main

import (
	"fmt"
)

var i int = 10
func main() {
	
	var (
		a int = 10  //1010
		b int = 3   //0011
		c int = 8 
	)
    //bit operations 
	fmt.Println("----first----")
	fmt.Println("BIT OPERATION   --   &  |  *  &^") 
	fmt.Println( a & b )  //0010
	fmt.Println( a | b )  //1011
	fmt.Println( a ^ b )  //1001
	fmt.Println( a &^ b ) //0100

	fmt.Println("----second----")
    fmt.Println("BIT SHIFT - Left 3 places") 
	// bit shift
	// bit shift c left 3 places, x * 2^y 
	// x << y = x * 2^y = 8 * 2^3 = 8 * 8 = 64
	fmt.Println(c << 3) 
	fmt.Println("----third----")
	fmt.Println("BIT SHIFT - Left 4 places")
	// bit shift c left 4 places, x * 2^y
	// x * 2^4 = 8 * 16 = 128
	fmt.Println(c << 4) 
	fmt.Println("----fourth----") 
	fmt.Println("BIT SHIFT - Right 3 places")
	// bit shift c right 3 places,
	// x >> y =  x / 2^y
	// 8 / 2^3 = 8/8 = 1
	fmt.Println(c >> 3)  
}