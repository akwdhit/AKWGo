// aURL: https://gobyexample.com/for
package GoByExample

import(
	"akw/akwutils"
	"fmt"
	"strconv"
)

func For_005() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)
	
	// 1. Single Condition loop (The most basic loop in Go)
	fmt.Println("These are the numbers that are printed by single condition loop: ")
	i := 0	// aStart or the incremental variable
	for (i < 3) {
		fmt.Println(i)
		// aCannot perform? ++i --> syntax error: unexpected ++, expecting }
		i++
	}
	
	// 2. Classic - "initial | condition | after" - for loop
	fmt.Println("These are the numbers started from 10 to 14")
	for j := 10; j < 15; j++ {
		fmt.Println(j)
	}
	
	// 3. For without condition. Will stop if:
	// 3a. there is a return statement of the enclosing function
	// 3b. there is a break statement of the loop
	k := 1000
	for {		
		if (k < 10) {
			// aOne way to print integer
			fmt.Println("Enough! " + strconv.Itoa(k))
			break
		} else {
			// aAnother way to print int
			fmt.Print("Current k: ")
			fmt.Printf("%d\n", k)
			k = k / 2
		}
	}
	
	// 4. Skipping loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		
		fmt.Println(n)
	}
}