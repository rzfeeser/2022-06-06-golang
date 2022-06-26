/* Alta3 Research | RZFeeser
   Errors - Returning a string and error from a function
   Errors are typically returned as the last argument in a function
   Errors are typically lower case and do not end in punctuation
   Errors should be returned as Nil if there is no error (this is the zero value) */

package main

import (
    "errors"
    "fmt"
)

// functions should return error as the "last" value
func obiwan(name string) (string, error) {
    if name == "" {
        return "", errors.New("no name provided")  // best practice all lower case
    }
    // the last position is the error
    return (name + ", now that is a name I haven't heard in a long time."), nil
}

func main() {

    var name string        // define a string variable (for tracking user name)

    fmt.Print("What was the name? ")  // "fmt.Print" is not followed by a /n
    fmt.Scanf("%s", &name)             // %s - string, and a pointer to "name" 
                                       // (otherwise we'll make a copy of name
                                       //  and change the copy)    

    // run the obiwan function
    name, err := obiwan(name)
    
    if err != nil {
        fmt.Println("Could not run the obiwan function:", err)
        return
    }

    // display what "Old Ben" Obi-Wan Kenobi has to say
    // ONLY if there has not been an error
    fmt.Println("Obi-Wan Says:", name)
}

