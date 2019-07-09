// aTo make it as util
package akwutils

// aDefining the correct imports
import (
	"errors"
	"runtime"
)

/* 
 * aCopied from https://stackoverflow.com/questions/47218715/is-it-possible-to-get-filename-where-code-is-called-in-golang
 * https://stackoverflow.com/posts/47219375/revisions
 * aWe should make the method exportable with an uppercase in front of its name: https://stackoverflow.com/questions/25501875/error-in-importing-custom-packages-in-go-lang
 */
func GetFileName() (string, int, string, error) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		/*
		 * aI don't want to print this out, instead I will send it
		 *
			fmt.Printf("Called from %s, line #%d, func: %v\n",
				file, line, runtime.FuncForPC(pc).Name())
		 */
		
		return file, line, runtime.FuncForPC(pc).Name(), nil
	}
	
	return "", 0, "", errors.New("Unhandled condition")
}