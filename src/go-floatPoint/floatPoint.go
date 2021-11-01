package main

import (
	"fmt"
)

func main() {
	fmt.Println("----first----")
	var n float64 = 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Println("PRINTING FLOATING POINT NUMBERS")
	//NOTE: %v=value %T=dataType
	fmt.Printf("%v %T\n", n, n)
	fmt.Println("----second----")
    fmt.Println("\nOPERATIONS ON FLOATING POINT NUMBERS")
	a := 10.2
	b := 3.7
	fmt.Println( a + b)
	fmt.Println( a - b)
	fmt.Println( a * b)
	fmt.Println( a / b)
	fmt.Println("----third----")
	fmt.Println("\nIMAGINARY NUMBERS")
	//Print imaginary number
	var m complex64 = 1 + 2i
	var x complex128 = 1 + 2i
	var w complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", m , m)
	fmt.Println("\nIMAGINARY NUMBERS SHOW ONLY REAL")
	fmt.Printf("%v, %T\n", real(m) , real(m))
	fmt.Println("\nIMAGINARY NUMBERS SHOW ONLY IMAGINARY")
	fmt.Printf("%v, %T\n", imag(m) , imag(m))
	fmt.Println("\nIMAGINARY NUMBERS SHOW ONLY REAL")
	fmt.Printf("%v, %T\n", real(x) , real(x))
	fmt.Println("\nIMAGINARY NUMBERS SHOW ONLY IMAGINARY")
	fmt.Printf("%v, %T\n", imag(x) , imag(x))
	fmt.Println("\n\nCOMPLEX NUMBERS")
	fmt.Printf("%v, %T\n", w, w)

	fmt.Println("----fourth----")
	fmt.Println("\nOPERATIONS ON IMAGINARY NUMBERS")
	//Print imaginary number
	c :=  1 + 2i
	d :=  2 + 5.2i
	fmt.Println( c + d)
	fmt.Println( c - d)
	fmt.Println( c * d)
	fmt.Println( c / d)
}