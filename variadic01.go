/* Alta3 Research | RZFeeser 
   Variadic Functions - Go Lang function that accepts an indefinite number
   of args */

package main

import "fmt"

// unpack the 'names' argument (effectively a slice)
// a "variadic" function accepts an indefinite number of a type
// the "..." must appear at the END of any other arguments passed to the function
func hello(names ...string) {
    
    // print any of the values passed in
    fmt.Print(names, " ")
    
    total := 0
    
    // the "blank identifier" is a place holder for the relative position
    // name is the "value" within names
    for _, name := range names {
        total += 1         // each time through the loop, increase total by 1
        fmt.Println("Hello", name) // say hello the to next person
    }
    
    fmt.Println(total)  // show the total
}


func main() {

    hello("larry", "steve")
    hello("jane", "raj", "tom")

    // create a slice of strings
    names := []string{"larry", "raj", "harry", "beth"}
    hello(names...)    // "unpack" our name slice
}

