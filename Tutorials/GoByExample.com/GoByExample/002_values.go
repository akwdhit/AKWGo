// aURL: https://gobyexample.com/values

// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

// aTo be able to be executed
package GoByExample

// aTo call priting methods
import (
	"akw/akwutils"
	"fmt"
)

// aMain method to be called
func Values_002() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)
	
	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")
	
	// aDefining string and cast as a rune (int32)
	str := "AKW Keren"
	fmt.Printf("%q\n", []rune(str))	// aCast string as rune
	fmt.Printf("%d\n", len(str))	// aLength of the rune casted from String

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Printf("7.0/3.0 = %d\n", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}