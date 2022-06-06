/* Alta3 Research | Author: RZFeeser
   Type transformations - string to integer  */

package main

import (
    "fmt"
    "strconv"    // built in package for string conversions
)

func main() {

    // create a string via inference
    s := "33"
    
    // string to int
    // see documentation @ https://pkg.go.dev/strconv#Atoi
    i, err := strconv.Atoi(s)
    
    // avoid nesting and check error immediately
    // less nesting means less load on the compiler
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("The type of s is %T and the value is %v\n", s, s)  // the type of s followed by the value of s
    fmt.Printf("The type of i is %T and the value is %v\n", i, i)  // the type of i followed by the value of i
            
    // transform an integer into a string
    // see documentation @ https://pkg.go.dev/strconv#Itoa
    st := strconv.Itoa(42)   // We have transformed 42 into a string
    fmt.Printf("The type of st is %T and the value is %v\n", st, st)  // the type of st followed by the value of st        
}

