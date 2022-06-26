/* Alta3 Research | RZFeeser
   for-each range loop */

package main

import "fmt"

func main() {
    
    // create a string slice
    a := []string{"Alta3", "Research", "Loves", "GoLang"}
    
    // the "i" position is always the index (position)
    // the "s" position is always the corresponding value
    // the second iteration variable ("s") is optional
    for i, s := range a {
        fmt.Println("Position", i, "contains the string:", s)
    }
    
    // in this example, we removed the second iteration variable ("s")
    // the "i" position is always the index (position)
    for i := range a {
        fmt.Println("Position", i)
    }

    // the name of the variable we use in the for-loop is not relevant
    // however, best practice says that all variables should be named in a
    // manner that is easy to understand
    for pos := range a {
        fmt.Println("Position", pos)
    }
    
    for _, val := range a {
        fmt.Println("We don't want the position, just the values", val)
    }


}

