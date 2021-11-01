package main

import (
	"fmt"
)

// Declaring an Anonomous Struct

func main() {
	//struct{name string} - defineds structure of struct
	//{name: "Don Dillon"} is initializer - provides data into struct
	aDoctor := struct{name string}{name: "Don Dillon"}
	fmt.Println(aDoctor)
    //struct create copy not refercence
	fmt.Println("\nNOTE: structs create copy not refercence")
	anotherDoctor := aDoctor
	anotherDoctor.name = "Xavier Dillon"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	fmt.Println("\nNOTE: can use POINTER refercence")
	myDoctor := &aDoctor
	myDoctor.name = "Xavier Dillon"
	fmt.Println(aDoctor)
	fmt.Println(myDoctor)
}