/* Alta3 Research | RZFeeser
   init - defining a variable within init */

package main

import "fmt"

var name string

func init() {
    fmt.Println("Think of the init function like a prelude to main.")
    name = "Alta3 Research"
}

func main() {
    fmt.Println("This would run 'after' the init.")
    fmt.Printf("Name: %s\n", name)     // this is a format string
}

