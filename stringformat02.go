/* Alta3 Research | RZFeeser
   String formatting - Indexing values */
package main

import (
    "fmt"
)

func main() {

    // create some int vars we can use for testing
    n2 := 2   // n2 is int 2
    n3 := 3   // n3 is int 3
    n4 := 4   // n4 is int 4

    // create the string, but don't render it yet
    res := fmt.Sprintf("There are %d oranges %d apples %d plums", n2, n3, n4)
    fmt.Println(res) // There are 2 oranges 3 apples 4 plums

    // Notice we not are using indexing
    // argument position 0 (the string)				      ,  1,  2,  3,     4
    res2 := fmt.Sprintf("There are %[2]d oranges %[3]d apples %[1]d plums", n2, n3, n4)
    fmt.Println(res2) // There are 3 oranges 4 apples 2 plums
}

