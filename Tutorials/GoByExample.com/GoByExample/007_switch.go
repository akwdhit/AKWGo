// aURL: https://gobyexample.com/switch
package GoByExample

import (
	"akw/akwutils"
	"fmt"
	"time"
)

// aMain method to be called in this file
func Switch_007() {
	// aTo know, from where it is called
	// _ here means ignoring that variable
	file, line, pcName, _ := akwutils.GetFileName()
	fmt.Printf("Called from %s, line #%d, func: %v\n", file, line, pcName)

	switchCase1()
	switchCase2()
	switchCase3()
	switchCase4()
	
}

// aCase 1: Use an object for switch
func switchCase1() {
	i := 2
	
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("eins")
	case 2:
		fmt.Println("zwei")
	case 3:
		fmt.Println("drei")
	}
	
	/*
	 * aSomehow the function equals to
	 
	 if i == 1 {
		fmt.Println("eins")
	 } else if i == 2 {
		fmt.Println("zwei")
	 } else if i == 3 {
		fmt.Println("drei")
	 }
	 */
}

// aCase 2: Use commas to separate multiple expressions in the same case statement
// aUse the optional default case as "else"
func switchCase2() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	
	/*
	 * aSomehow the function equals to
	 
	 day := time.Now().Weekday()
	 if day == time.Saturday || day = time.Sunday {
		fmt.Println("It's the weekend")
	 } else {
		fmt.Println("It's a weekday")
	 }
	 */
}

// aCase 3: switch without an expression (alternate way to define if/else logic).
// aAlso to show that the express could contain non-constants
func switchCase3() {
	t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}

// aCase 4: Type switch - to compare types instead of values (use to discover the type of an interface)
func switchCase4() {
	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Unknown type, let me ask the interpretor: %T\n", t)
        }
    }
	
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}