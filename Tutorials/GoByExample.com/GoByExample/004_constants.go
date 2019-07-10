// aURL: https://gobyexample.com/constants
package GoByExample

import (
	"akw/akwutils"
	"fmt"
	"math"
)

// aForget public static final String temporarily in Go, use this instead! :( But, I've missed that already :(
// const [variable name] [variable type] = [declarated value]
const s string = "constant"
const i int = 1

// aDon't forget to make it exportable by giving an uppercase in front
func Constants_004() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)
	
	fmt.Println(s)
	
	fmt.Println(i)
	
	// A const statement can appear anywhere a var statement can.
    const n = 500000000
	
	// Constant expressions perform arithmetic with arbitrary precision.
    const d = 3e20 / n
    fmt.Println(d)
	
	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
    fmt.Println(int64(d))
	
	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
    fmt.Println(math.Sin(n))
}