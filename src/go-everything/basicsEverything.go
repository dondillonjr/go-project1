package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name    string
	age     int
	address string
}

//func TestAddition(t *testing.T) {
//if Addition(2 , 3) != 5 {
//t.Error("Expected 2 + 3 to equal 5")
////}
//}

func main() {
	var name = "DON"
	var x int = 2
	var y int = 3
	var total int = x + y
	fmt.Println("-----first-----")
	fmt.Println(" 2 + 3=", total)
	fmt.Println("-----second-----")
	//go can infer type, short hand type
	z := 5
	w := 3
	t := z + w
	fmt.Printf("z + w=%v\n", t)
	fmt.Println("Use v for value, T for data-type")
	fmt.Printf("%v, %T\n", z, z)
	fmt.Println("HI ")
	fmt.Println("-----third-----")
	//array - fixed length
	var a [5]int
	a[0] = 7
	fmt.Println("array index[0]=", a[0])
	fmt.Println("-----fourth-----")
	b := [5]int{1, 2, 3, 4}
	fmt.Println("array b members=", b)
	fmt.Println("-----fifth-----")
	// slice like an array but allows size increase via append
	// has no index in declaration
	c := []int{1, 2, 3, 4, 5}
	c = append(c, 6)
	fmt.Println("array c holds: ", c)
	fmt.Println("-----sixth-----")
	// Maps  has key and value, must use "make" function
	calender := make(map[string]int)
	calender["January"] = 1
	calender["Febuary"] = 2
	calender["March"] = 3
	calender["April"] = 4
	calender["May"] = 5
	calender["June"] = 6
	calender["July"] = 7
	calender["August"] = 8
	calender["September"] = 9
	calender["October"] = 10
	calender["November"] = 11
	calender["December"] = 12

	fmt.Println("HELLO", name)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("December has: ", calender["December"], "days")
	fmt.Println(calender)
	fmt.Println("-----seventh-----")
	delete(calender, "December")
	fmt.Println(calender)

	fmt.Println("-----eighth-----")
	//go only uses 'for' loops no while
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----ninth-----")
	// can act as while loop
	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
	fmt.Print("\n")

	fmt.Println("-----tenth-----")
	h := []int{1, 2, 3, 4, 5, 6, 7}
	// transverse through a loop by using range in for loop
	for index, value := range h {
		fmt.Println("index=", index, " value=", value)
	}

	fmt.Println("-----eleventh-----")
	// transverse through a map
	for key, value := range calender {
		fmt.Println("key=", key, " value=", value)
	}

	fmt.Print("\nCalculate :", sum(5, 5))

	fmt.Println("-----twelfth-----")
	result, err := sqrt(16)
	//if, else if, else
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nSQRT= ", result)
	}

	fmt.Println("-----thirteenth-----")
	result2, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nSQROOT= ", result2)
	}

	fmt.Println("-----fourteenth-----")
	p := person{name: "Don Dillon", age: 50, address: "340 MillHaven"}
	fmt.Println()
	fmt.Println(p)
	fmt.Println(p.name)

	fmt.Println("-----fifteenth-----")
	var num = 60
	fmt.Println(&num)

	fmt.Println("-----sixteenth-----")
	inc_1(num)
	fmt.Println("num in main after inc_1()=", num)
	fmt.Println()
	inc_2(&num)
	fmt.Println()
	fmt.Println("num in main after inc_2()=", num)
}

//function with return value
func sum(x int, y int) int {
	return x + y
}

/* function can have multiple return values
   go does not have exception
*/
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("\nUndefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

// inc function gets a copy of varable x - called copy by value
// not original varable x is not modified
func inc_1(x int) {
	x++
	fmt.Println("value of X inside inc_1()=", x)
}

/* function looks at varable x memory reference address, original varable X
   can be modified. If given the address you must de-reference to access value
*/
func inc_2(x *int) {
	fmt.Println("I am in func inc_2: ", x)
	fmt.Println("func inc_2, value of X before increment: ", *x)
	*x++
	fmt.Println("func inc_2(), value of X after increment:", *x)
}
