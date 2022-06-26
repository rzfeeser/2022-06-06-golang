/* Alta3 Research | RZFeeser
   Maps - Using empty interfaces as a "wildcard" to create
   composite types */

package main

import "fmt"


func main() {

    // instead of declaring a type, we have supplied the "empty interface"
    // interface{} be in the place of the key, value, or both when making the map
    //futurama := make(map[string]interface{})
    futurama := make(map[string]any)

    // now we can create a map of mixed types
    futurama["name"] = "Turanga Leela"  // string: string
    futurama["age"] = 30                // string: int
    futurama["height"] = 182.5          // string: float

    // display the "mixed" results
    fmt.Printf("%+v\n", futurama)
}

