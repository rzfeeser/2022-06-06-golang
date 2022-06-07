/* RZFeeser | Alta3 Research
   Arrays - selecting value by index   */

package main

import "fmt"

func main() {
    var myArray [3]string
    myArray[0] = "United States"  // Assign a value to the first element
    myArray[1] = "Canada"         // Assign a value to the second element
    myArray[2] = "Japan"          // Assign a value to the third element

    fmt.Println(myArray)          // Display the entire myArray

    fmt.Println(myArray[0])       // Access the first element value
    fmt.Println(myArray[1])       // Access the second element value
    fmt.Println(myArray[2])       // Access the third element value
}

