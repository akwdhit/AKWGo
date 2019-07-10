// aURL: https://gobyexample.com/arrays
package GoByExample

import (
	"akw/akwutils"
	"fmt"
)

// aMain method to be called by the library / util
func Arrays_008() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)
	
	arrayLogic01()
	arrayLogic02()
	arrayLogic03()
}

// aFirst logics for the array
func arrayLogic01() {

	// aWhen an array is defined, it is by default has its own zero-valued
	var arr [5]int
	fmt.Println("First defined arr: ", arr);
	
	// aTry to set any index of the array, in this example index = 4
	arr[4] = 100
	fmt.Println("After set, the whole arr is: ", arr)
	fmt.Println("Get index 4: ", arr[4])
	
	// aLength of the arr
	fmt.Println("len: ", len(arr))
	
}

func arrayLogic02() {
	// aDeclare and initialize syntax in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b: ", b)
	
	// aDefining syntax for string
	c := [3]string{"eins", "zwei", "drei"}
	fmt.Println("c: ", c)
	
	// aAppending (adding at the tail) of the slice
	var e []string
	d := append(e, "eins", "zwei", "drei", "vier", "funf")
	fmt.Println("d: ", d)
}

func arrayLogic03() {
	var twoDimension [2][3]int
	
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimension[i][j] = i + j
		}
	}
	
	fmt.Println("twoDimension: ", twoDimension)
}