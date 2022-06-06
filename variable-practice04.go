/* Alta3 Research | RZFeeser
   Variables - Short variable declaration */

package main

import (
    "fmt"
)

// t is a package-level variable that can be accessed
// anywhere in the `main` package. You can say that it is
// global to the package
var t = true

func main() {
    // f is a block-level variable accessible from its point of declaration to
    // to the end of the function
    f := false

    {
    // i is a block-level variable that's only valid from this point
    // until the end of the block
        i := 20
        fmt.Println(i) // 20
    } // this block scope is now over so i is no longer valid

    fmt.Println(t, f) // true false
}

