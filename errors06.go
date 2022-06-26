/* Alta3 Research | RZFeeser
   Predefined Errors - Defining expected errors,
   to explicitly check for this error using "errors.Is" */

package main

import (
    "errors"
    "fmt"
)

// Pre-define an expected error
var ErrDivideByZero = errors.New("divide by zero")

// a simple division function that may encounter the divide by zero
func Divide(a, b int) (int, error) {
    // if b is passed as 0, return the ErrDivideByZero
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil  // the operation we want to perform
}

func main() {
    a, b := 10, 0      // create an intentional error
    result, err := Divide(a, b)    // pass 10 and 0 for b which will trigger our error
    if err != nil {
        switch {
        // this case will match "true" if the error matches this condition
        case errors.Is(err, ErrDivideByZero):
            fmt.Println("divide by zero error")
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}

