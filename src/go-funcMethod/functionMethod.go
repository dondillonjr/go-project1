package main

import (
	"fmt"
)

type counter int

type greeter struct {
	greeting string
	name string
}

//Method Function
//(g greeter) makes function a METHOD
//A METHOD is a function that is executing in a known 
//context, i.e, any type, struct, int
// 
//greet() METHOD is going to get a copy of the 
//greeter struct object (g), in the context that function
//is executing in.
//Note: (g greeter) is known as a value reciever  
//We are getting a copy of the struct
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	//NOTE
	g.name = ""  //effects copy, not original struct
}

//Note: (g greeter) here is known as a pointer reciever
//pointerGreet() METHOD is going to get access to the 
//original (greeter) struct object (g), in the context that function
//is executing in.
func (g *greeter) pointerGreet() {
	//implicate de-referencing of pointers here.
	fmt.Println(g.greeting, g.name)
	//NOTE
	g.name = ""  //effects original struct
}

func (g counter) doCount() {
	var x counter = 0
	for  x < g {
		fmt.Println(x)
		x++
	}
}

func main() {
	//initialize struct variable (g) with fields values
	g := greeter {
		greeting: "Hello", 
		name: "Dillon",
	}
	fmt.Println("----first----")
	//Syntax for calling Method greet(), preceeding it with struct (g)
	g.greet()
	fmt.Println("The name is still: [", g.name,"]")

	fmt.Println("----second----")
	//Syntax for calling Method pointerGreet(), preceeding it with pointer to struct (g)
	g.pointerGreet()
	fmt.Println("The name is now:[", g.name,"]")

	fmt.Println("----third----")
	var c counter = 8
	c.doCount()
	fmt.Println("Final Count=", c)
}