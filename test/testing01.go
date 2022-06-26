/* Alta3 Research | RZFeeser
   Testing - file on which to run a simple test */
   
package main

import (
    "fmt"
    "errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf("Hello there %v", name)
    return message, nil
}

