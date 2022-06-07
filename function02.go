/* RZFeeser | Alta3 Research
   nil in the error position means "no error" */
   
package main

import (
    "fmt"
)

//name of funct   //input                      //output
func rollchar(firstName string, lastName string) (string, error) {
    return firstName + " the " + lastName, nil
}

func main() {
    fmt.Println("Welcome to the Character Generator")

    playerChar, err := rollchar("Gandalf", "Grey")

    if err != nil {
        fmt.Println("Error while spawning your requested character.")
    } else {
        fmt.Println(playerChar, "has been generated.")
    }
}

