/* Alta3 Research | Author: RZFeeser
  Types - transformations with buildin functions. Review the standard library @
    https://pkg.go.dev/std
      and
    https://pkg.go.dev/builtin@go1.18 */

package main

import (
    "fmt"
    "math"
)

func main() {
    
    // x and y are integers
    var earth, mars int = 3, 4
    fmt.Printf("The type of earth is %T and the value is %v\n", earth, earth)
    fmt.Printf("The type of mars is %T and the value is %v\n", mars, mars)
    
    // See documentation @ https://pkg.go.dev/math#Sqrt
                            // transform x and y into a float64()
                            // See documentation @ https://pkg.go.dev/std
                            // func float64() defined @ https://pkg.go.dev/builtin@go1.18
    var f float64 = math.Sqrt(float64(earth*earth + mars*mars))
    fmt.Printf("The type of f is %T and the value is %v\n", f, f)
    
    // tranform into unsigned integer
    var zebra uint = uint(f)  // we are passing by assignment
    fmt.Printf("The type of zebra is %T and the value is %v\n", zebra, zebra)
    
    // All of our operations have been passing by value (not by pointer -- we can cover this later)

    // IncDec statements
    // See documentation @ https://go.dev/ref/spec#IncDec_statements
    // add one to earth
    earth++ // increase by the "untyped constant" 1
    fmt.Printf("The type of earth has increased by one. The type is still %T and the value is now %v\n", earth, earth)

    // remove one from mars
    mars--  // decrease by the "untyped constant" 1
    fmt.Printf("The type of mars has decreased by one. The type is still %T and the value is now %v\n", mars, mars)

}

