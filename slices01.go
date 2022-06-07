/* RZFeeser | Alta3 Research
   Slices relationship to arrays     */

package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}   // this array will back the slice
    slice := array[:]                // [:] means "from start to the end of"

    //Modifying the slice (which is a pointer to the array)
    slice[1] = 7
    fmt.Println("Modifying Slice")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)

    //Modifying the array. Would reflect back in slice too
    array[1] = 2
    fmt.Println("Modifying Underlying Array")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)
}

