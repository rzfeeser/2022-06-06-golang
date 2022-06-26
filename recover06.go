/* Alta3 Research | RZFeeser
   Refer - Greenskeeper routine with panic(), defer, and recover() */

package main

import (
    "fmt"
)

func waterLawn() {
    defer func() { // this must be named func()
        if r := recover(); r != nil {
            fmt.Println("Error detected and recovered in the watering system (waterLawn):", r)
        }
    }()

    fmt.Println("Begin watering system.")

    water(5) // we will imagine that there is a bug in our program
             // a value of 3 or higher will crash our program
    
    
    
    fmt.Println("Watering system finished.")
}

func water(g int) {
    fmt.Println("Starting", g, "watering systems.")

    for i := 1; i < g+1; i++ {
        fmt.Println("Turning on sprinkler at station: ", i)
        defer fmt.Println("Turning off sprinkler at station: ", i)

        // this is the imaginary "bug" in our function
        // it only happens if i becomes equal to 3
        if i == 3 {
            fmt.Println("Panicking!")
            // panic requires input of a single sting
            panic(fmt.Sprintf("Error occurred on iteration %v with function input %v", i, g))
        }
    }
}

func main() {

    waterLawn()

    // two imaginary functions we might still want to have run
    // even through waterLawn() recovers from a failue
    // measureSoilPH()
    // measureSoilTemp()
    fmt.Println("Begin measuring soil pH.")
    fmt.Println("Begin measuring soil temperature.")

    fmt.Println("Completed Golf Course Greenskeeper Routine")
}

