package main

import (
	"fmt"
)


// Arrays can be declared as follows:
// a := [3]int{1,2,3}
// a := [...]int{1,2,3}
// var a [3]int
//
// Slices can be declared as follows:
// v := []int{1,2,3,4,5,6,7,8,9,10}
// u := [...]int{1,2,3,4,5,6,7,8,9,10}
// h := make([]int, 3)
// i := make([]int, 3, 20)  //has 3 elements but 20 cells
func main() {
	//grades := [3] int {97, 85, 93}
	fmt.Println("----first----")
	fmt.Print("Arrays\n")
	grades := [...]int {97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Println("----second----")
	var students [3]string
	fmt.Printf("\nStudents: %v\n", students)
	students[0] = "DON"
	students[1] = "Bolden"
	students[2] = "Dillon"	
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students #2: %v\n", students[2])
	fmt.Printf("Number of Students: %v\n", len(students))
	fmt.Print("\nMulti Dimentional Arrays\n")
	//var indentityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}
	//    OR
	// Declares an array of 3 arrays
	//   [[1 0 0] [0 1 0] [0 0 1]]
	fmt.Println("----third----")
	fmt.Println("Multi Dementional Array")
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1,0,0}
	identityMatrix[1] = [3]int{0,1,0}
	identityMatrix[2] = [3]int{0,0,1}
	fmt.Println(identityMatrix)
	fmt.Println("----fourth----")
    fmt.Print("\nCopy of an Array\n")
	a := [...]int{1,2,3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("----fifth----")
	fmt.Print("\nReference to an Array (pointer)")
	c := [...]int{1,2,3}
	d := &c
	d[1] = 7
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("----sixth----")
	fmt.Print("\nSlices are like Arrays - are like References\n")
	// Slice
	// Slices are references
	f := []int{1,2,3}
	g := f
	g[1] = 2
	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))
	fmt.Printf("Length: %v\n", len(g))
	fmt.Printf("Capacity: %v\n", cap(g))
	fmt.Println("----seventh----")
	// Copy elements of a Slice to another Slice
	fmt.Printf("\nCopy elements of a Slice to another Slice\n")
	v := []int{1,2,3,4,5,6,7,8,9,10}
	u := [...]int{1,2,3,4,5,6,7,8,9,10}
	w := v[:]   // copy all elements in v to w
	x := v[3:]  // slice from 4th element to end
	y := v[:6]  // slice first 6 elements
	z := v[3:6] // slice 4th, 5th, and 6th element
	v[5] = 42
	fmt.Println(v)
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(u)
	fmt.Println("----eight----")
	//use make to create a slice
	//Slices do not have to have fixed size
	fmt.Println("\nUse Make to create slice")
    h := make([]int, 3)
	i := make([]int, 3, 100)  //has 3 elements but 20 cells
	j := []int{}
	fmt.Println(h)
 	fmt.Printf("LengthFor-H: %v\n", len(h))
	fmt.Printf("CapacityFor-H: %v\n\n", cap(h))
	fmt.Println(i)
	fmt.Printf("LengthFor-I: %v\n", len(i))
	fmt.Printf("CapacityFor-I: %v\n\n", cap(i))
	fmt.Println(j)
	fmt.Printf("LengthFor-J: %v\n", len(j))
	fmt.Printf("CapacityFor-J: %v\n", cap(j))
	fmt.Println("----ninth----")
	//Add element to slice
	fmt.Println("\nAdd element to slice")
	j = append(j, 1)
	fmt.Printf("j=: %v\n ", j)
	fmt.Printf("LengthFor-J: %v\n", len(j))
	fmt.Printf("CapacityFor-J: %v\n", cap(j))
	fmt.Println("----tenth----")
	fmt.Println("\nAdd elements to slice")
	j = append(j, 2,3,4,5,6) //variatic function
	fmt.Printf("j=: %v\n ", j)
	fmt.Printf("LengthFor-J: %v\n", len(j))
	fmt.Printf("CapacityFor-J: %v\n", cap(j))
	fmt.Println("\nConcat slices")
	fmt.Println("----eleventh----")
	j = append(j, []int{7,8,9,10}...)
	fmt.Printf("j=: %v\n ", j)
	fmt.Printf("LengthFor-J: %v\n", len(j))
	fmt.Printf("CapacityFor-J: %v\n", cap(j)) 
	fmt.Println("----twelveth----")
	fmt.Println("\nStack Operations with slices")
	//append acts as push, shift as pop
	// Pop off at front
	k := j[1:]  //pops off element at index Zero and before
	fmt.Println("pops off element at index Zero and before")
	fmt.Printf("K=%v\n", k)
	fmt.Printf("J=%v\n", j)
	fmt.Println("----thirteenth----")
	l := j[2:]  //pops off element at index One and before
	fmt.Println("\npops off element at index One and before")
	fmt.Printf("L=%v\n", l)
	fmt.Printf("J=%v\n", j)
	fmt.Println("----fourteenth----")
	m := j[:len(j) -1]  //pops off element at index Nine 
	// Pop off at end
	fmt.Println("\npops off element at end")
	fmt.Printf("M=%v\n", m)
	fmt.Printf("J=%v\n", j)
	fmt.Printf("LengthFor-J: %v\n", len(j))
	fmt.Println("----fifteenth----")
	// Pop off elements after index One
	n := append(j[:2])
	fmt.Println("\npops off elements after index One")
	fmt.Printf("N=%v\n",n)
	fmt.Printf("J=%v\n", j)
	fmt.Println("----sixteenth----")
	// Pop off elements after index One
	o := append(j[5:])
	fmt.Println("\npops off elements after index One")
	fmt.Printf("O=%v\n",o)
	fmt.Printf("J=%v\n", j)
	fmt.Println("----seventeenth----")
	//Pop off elements in middle of slice
	//pop off element after index 1
	//pop off element at index 4 and before
	fmt.Println("\npops off elements in middle of slice")
	//j[:2]= 1,2
	//j[5:]= 6,7,8,9,10
	fmt.Printf("J=%v\n", j)
	p := append(j[:2], j[5:]...)
	fmt.Printf("J=%v\n", j)
	fmt.Printf("P=%v\n", p)	
}