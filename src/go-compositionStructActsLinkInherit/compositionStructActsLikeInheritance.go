package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	// Embeded Animal
	Animal
	SpeedKPH  float32
	CanFly    bool
}
//  Embed Animal Struct into Bird Struct
func main() {
	fmt.Println("---first---")
    b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 49
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)

	fmt.Println("---second---")
	// Can be done also as
	c := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH:  48,
		CanFly:    false,
	}
	fmt.Println(c.Origin)
}