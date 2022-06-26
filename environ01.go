/* Alta3 Research | RZFeeser
   Environmental Variables - Setting and reading environmental variables
   from an environment */
   
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    // set the environmental variable "ALTA" to "3"
    os.Setenv("ALTA", "3")
    fmt.Println("ALTA:", os.Getenv("ALTA"))
    
    // we did not set this value
    fmt.Println("RESEARCH:", os.Getenv("RESEARCH"))

    // print a blank line
    fmt.Println()
    
    // loop across all of the environmental variables on the system
    // throw out the positions and only take the value
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)    // split "e" across the "="
        fmt.Println(pair[0]) 
    }
}

