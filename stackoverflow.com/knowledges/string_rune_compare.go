// aVon: https://stackoverflow.com/questions/19310700/what-is-a-rune
// aOne thing that I learn at least for now, is that, if you count this string + chinese chars as 7, then you need to convert it as rune instead of using it as a string.
// aAs they always said, go see String as slice of bytes (can't understand this term yet, hopefully soon)

// string & rune compare,
// aMake it package main to be executable as a command
package main

// aImport the library to print
import (
	"fmt"
)

// string & rune compare,
func stringAndRuneCompare() {
    // string,
    s := "hello你好"

    fmt.Printf("%s, type: %T, len: %d\n", s, s, len(s))
    fmt.Printf("s[%d]: %v, type: %T\n", 0, s[0], s[0])
    li := len(s) - 1 // last index,
    fmt.Printf("s[%d]: %v, type: %T\n\n", li, s[li], s[li])

	// aCast the string as []rune (So, that each character in the string (utf-8) is converted to codepoint (int32))
    // []rune
    rs := []rune(s)
    fmt.Printf("%v, type: %T, len: %d\n", rs, rs, len(rs))
}

// aThe main method to be called
func main() {
    stringAndRuneCompare()
}