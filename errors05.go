/* Alta3 Research | RZFeeser
   Errors - Throwing out results with _ */

package main

import (
    "errors"
    "fmt"
)

func obiwan(name string) (string, error) {
    if name == "" {
        return "", errors.New("no name provided")  // best practice all lower case
    }
    return (name + ", now that is a name I haven't heard in a long time."), nil
}

func main() {

    var name string        // define a string variable (for tracking user name)

    fmt.Print("What was the name? (Just press enter)") 
    fmt.Scanf("%s", &name)             // %s - string, and a pointer to "name" 
                                       // (otherwise we'll make a copy of name
                                       //  and change the copy)    

    // run the obiwan function
    // use of the "_" means "throw out these results"
    _, err := obiwan(name)
    
    // without capturing the results of obiwan()
    // all we can do is error check
    if err != nil {
        fmt.Println("Could not run the obiwan function:", err)
        return
    }
}

