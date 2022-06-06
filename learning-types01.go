/* Alta3 Research | Author: RZFeeser
   Types - learning about Go types   */

package main

import (
    "fmt"
    "math/cmplx"
)


// variables declarations may be "factored"
// into blocks (as with import statements)

var (
    // bool - boolean is true or false (lower case)
    kirk   bool       = false
    
    // uint64 - unsigned integers with values
    //          ranging from 0 to 18,446,744,073,709,551,615
    spock  uint64     = 1<<64 - 1               // Create a huge number by shifting a 1 bit left 64 places
                                                // and then subtracting ONE
    
    // complex 64 and complex128 types
    // Complex numbers consist of "real" and "imaginary" parts
    // Documentation avail @ https://pkg.go.dev/math/cmplx#Sqrt
    mccoy  complex128 = cmplx.Sqrt(-5 + 12i)    // -5 is the "real part"
                                                // 12i is the "imaginary part"
)

func main() {
    // %T reveals the "type" whereas %v matches the type to a string, \n is just a newline
    fmt.Printf("Type: %T Value: %v\n", kirk, kirk)
    fmt.Printf("Type: %T Value: %v\n", spock, spock)
    fmt.Printf("Type: %T Value: %v\n", mccoy, mccoy)
}

