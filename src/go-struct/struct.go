package main

import (
	"fmt"
)

//  Naming convention for Variable
// Capital - var "struct Doctor" is to be exported from package
// lowerCase - fields " number, actorName, familyMembers is internal to package main
// Other packages could see struct Doctor by nothing in any other package would see fields
// To allow other packages to see/use struct Doctor and its fields both struct name and field names must begin with capital letter.
type Doctor struct {
	number int
	actorName string
	familyMembers []string
}

func main() {
	aDoctor := Doctor{
		number: 3,
		actorName: "Don Dillon",
		//[]string is a slice
		familyMembers: []string {
			"Rhonda Dillon",
			"Donnie Dillon",
			"Xavier Dillon",
			"Etienne Dillon", 
		},
	}
	//Print all items in struct
	fmt.Println("Print all items in struct")
    fmt.Println(aDoctor)
	fmt.Println("\nPrint actorName")
	fmt.Println(aDoctor.actorName)
	fmt.Println("\nPrint familyMembers")
	fmt.Println(aDoctor.familyMembers)
	fmt.Println("\nPrint familyMember=[Xavier]")
	fmt.Println(aDoctor.familyMembers[2])
}