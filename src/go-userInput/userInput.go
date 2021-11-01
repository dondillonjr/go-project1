package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var fname string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	fname, _ := reader.ReadString('\n')

	newreader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a decimal number: ")
	number, _ := newreader.ReadString('\n')
	pnumber, _ := strconv.ParseFloat(strings.TrimSpace(number), 64)

	fmt.Print("Your Full Name is " + fname)
	fmt.Print(pnumber + 2)
}
