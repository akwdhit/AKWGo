// aURL: https://gobyexample.com/variables
package GoByExample

import (
	"akw/akwutils"
	"fmt"
)

func Variables_003() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)
	
	var a = "initial"
    fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
    fmt.Println(b, c)

	// Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    var e int
    fmt.Println(e)

    // The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
    f := "apple"
    fmt.Println(f)
}