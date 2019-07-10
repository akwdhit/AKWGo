// aURL: https://gobyexample.com/if-else
package GoByExample

// aImport statements
import (
	"akw/akwutils"
	"fmt"
)

// aMain method to be called for this library/util
func If_006() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)
	
	// aIf could have no else, could have an else, or elses
	// aA statement can precede conditionals, any variables 
	//  declared in this statement are available in all branches
	if num := getValue(); num < 0 {
		fmt.Println("Num is negative")
	} else if num < 10 {
		fmt.Println("Num has 1 digit")
	} else {
		fmt.Println("Num has multiple digits")
	}
	
	/*
	 * Note for IF
	 * -----------
	 * 1. You don't need parentheses in the condition, but braces are required
	 * 2. There is no ternary if in Go ( (condition)? statementTrue : statementFalse )
	 */
}

// aSimple method just to return an integer
func getValue() int {
	return 1;
}
