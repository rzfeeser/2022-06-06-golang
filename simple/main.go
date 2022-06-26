/* Alta3 Research | RZFeeser
   Simple GoLang program      */

// define the package name in which the program lie
// mandatory statement, as Go programs run in packages
package main

// include these files
import (
    "fmt"    // include files in the "fmt" package
)

// This is where the program execution begins
func main() {
    fmt.Println("I am IronMan.")     // display to standard out
    
    // Notice the capital P of Println method
    // in Go a name is exported if it started with a capital letter
    // Exported means the function or variable/constant is accessible to the
    // importer of the respective package 
}

