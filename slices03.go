/* RZFeeser | Alta3 Research
   Exploring slices - length vs capacity  */

package main

import "fmt"

func main() {

    // create a length of 3 and capacity of 5
    numbers := make([]int, 3, 5)
    fmt.Println("numbers =", numbers)         // 0 0 0 nil nil (nil values not shown)
    fmt.Println("length =", len(numbers))     // 3
    fmt.Println("capacity =", cap(numbers))   // 5

    // This line will cause a runtime error index out of range [4] with length 3
    //numbers[4] = 5

    //Increasing the length from 3 to 5
    numbers = numbers[0:5]  // 0 0 0 0 0

    // This line will no longer cause a runtime error
    numbers[4] = 5   // 0 0 0 0 5

    fmt.Println("Increasing length from 3 to 5")
    fmt.Println("numbers =", numbers)          // 0 0 0 0 5
    fmt.Println("length =", len(numbers))      // 5
    fmt.Println("capacity =", cap(numbers))    // 5

    // just use append in most situations
    numbers = append(numbers, 66, 777, 8888)   // 0 0 0 0 5 66 777 8888
    // numbers = append(numbers, "cake")
    fmt.Println("numbers =", numbers)
}

