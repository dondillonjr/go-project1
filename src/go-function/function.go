package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----first-----")
	sayHello("HELLO DON")
	fmt.Println("----")
	for i :=0; i < 5 ; i++ {
		sayMessage("SAY HELLO DON", i)
	}
	fmt.Println("---second----")
	sayHy("Hello", "Rhonda")

	fmt.Println("---third----")
	sayHy2("Hello", "Xavier")

	fmt.Println("---fourth----")
	greeting := "Hy"
	name := "Etienne"
	sayHy3(greeting, name)
	fmt.Println(name)

	fmt.Println("---fifth----")
	greetings := "Hy"
	names := "Etienne"
	sayHy4(&greetings, &names)
	fmt.Println(names)

	fmt.Println("---sixth----")
	greet := "Hy"
	myName := "Etienne"
	sayHy4(&greet, &myName)
	fmt.Println(myName)
}

func sayHello(msg string) {
	fmt.Println(msg)
}

func sayMessage(msg string, idx int) {
	 fmt.Println(msg)
	 fmt.Println("The value of the index is=", idx)
}

func sayHy(greeting string, name string) {
	fmt.Println(greeting, name)
}

func sayHy2(greeting, name string) {
	fmt.Println(greeting, name)
}

func sayHy3(greeting, name string) {
	name = "Donnie"
	fmt.Println(name)
}

func sayHy4(greeting *string, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Donnie"
	fmt.Println(*name)
}

func sayHy5(greeting, name *string) {
	fmt.Println(greeting, *name)
	*name = "Donnie"
	fmt.Println(*name)
}