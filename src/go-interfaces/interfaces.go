package main

import (
	"fmt"
) 

// Interfaces are a type
// type keyword, name of interface, type created is type interface
// surround definition with currly braces
//
// Interfaces describe behavior - will store method definitions
// Interface defined below:
type Writer interface {
	//except a slice of bytes, returns an int and error
	//Whatever implaments this interface, will take in
	//a slice of bytes, write it to something. Integer is
	//the number of bytes writen
	Write([]byte) (int, error)
}

//Implement Interface (Writer) with a (ConsoleWriter)
//implementation - see below:
type ConsoleWriter struct {}

//Create a method on ConsoleWriter that has the signiture
//of a Writer interface
//
//Below: METHOD called Write, excepts slice bytes,
//returns an int (# of bytes writen) or error.
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	fmt.Println("-----first-----")
	// Initialize var (w) is of type Writer which is set equal to
	// a ConsoleWriter instance
	// (w) variable is holding a writer which is something that 
	// implaments the (Writer) interface
	var w Writer = ConsoleWriter{}

	//Invoke Write Method 
	w.Write([]byte("Hello Don!"))
}