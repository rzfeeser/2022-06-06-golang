/* Alta3 Research | RZFeeser
   Errors - Returning errors from functions */

package main

import (
    "errors"
    "fmt"
)

// notice that error is the type being returned
// the universal block describes this as being a basic type
// accessible to all programs (along with int, string, true, false, etc.)
func endit() error {
    // this function always returns an error
    return errors.New("we don't have the power")
}

func main() {

    // call the function endit that returns na error
    err := endit()

    // we will always find an error (thats all that endit() can produce)
    if err != nil {                             // this paradigm replaces try/catch from other languages
    fmt.Println("We detected and error:", err)
        return  // return from main()
    }
    
    // this will only run if no error occurs
    fmt.Println("Release the hounds!")
}

