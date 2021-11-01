package main

import (
	"fmt"
)

//Note: Slices and MAPs in go act like references
func main() {
	a := map[string]string {"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a , b)
	a["foo"] = "qux"
	//Note (b) the value for foo is also (qux)
	fmt.Println(a , b)
}