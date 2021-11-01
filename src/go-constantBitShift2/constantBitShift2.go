package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota // 1 * 2^0  00000001
	isHeadquarters      // 1 * 2^1
	canSeeFinacials     // 1 * 2^2  100

	canSeeAfrica        // 1 * 2^3
	canSeeAsia          // 1 * 2^4
	canSeeEurope        // 1 * 2^5 100000
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	fmt.Println("----first----")
	var rr byte = isAdmin
	fmt.Printf("%v %b\n", isAdmin, rr)
	rr  = isHeadquarters
	fmt.Printf("%v %b\n", isHeadquarters, rr)
	rr  = canSeeFinacials
	fmt.Printf("%v %b\n", canSeeFinacials, rr)
	rr  = canSeeAfrica
	fmt.Printf("%v %b\n", canSeeAfrica, rr)
	rr  = canSeeAsia
	fmt.Printf("%v %b\n", canSeeAsia, rr)
	rr  = canSeeEurope
	fmt.Printf("%v %b\n", canSeeEurope, rr)

	fmt.Println("----second----")
	var roles byte = isAdmin | canSeeFinacials | canSeeEurope
	fmt.Printf("\n%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is HQ? %v", isHeadquarters & roles == isHeadquarters)
}