package main

import ( 
	"fmt"
)
 
const (
	//// bit shift (10 * iota), x * 2^y 
	_ = iota // ignore first value by assigning to blank 
	KB = 1 << (10 * iota)	// 1 * 2^10 = 1024
	MB  // 2 * (10 * 2)     // 1 * 2^20 = 1048576
	GB  // 2 * (10 * 3)     // 1 * 2^30 = 1073741824
	TB  // 2 * (10 * 4)     // 1 * 2^40 = 1099511627776
	PB  // 2 * (10 * 5)
    EB  // 2 * (10 * 6)
	ZB  // 2 * (10 * 7)
	YB  // 2 * (10 * 8)
)

func main() {
	//fileSize := 4000000000.
	fileSize := 4000000000.
	fmt. Printf("%v %.2f KB\n", KB, fileSize/KB)
	fmt. Printf("%v %.2f MB\n", MB, fileSize/MB)
	fmt. Printf("%v %.2f GB\n", GB, fileSize/GB)
	fmt. Printf("%v %.2f TB\n", TB, fileSize/TB)	
}