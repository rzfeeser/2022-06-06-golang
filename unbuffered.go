/* Alta3 Research | RZFeeser
   Channels - A simple example  */

package main

import (
    "fmt"
    "math/rand"
    "time"
)


// this wants a channel to be passed to it
// think of a channel like a 'wormhole'
func CalculateValue(values chan int) {
    value := rand.Intn(10)  // create a random value
    fmt.Println("Calculated Random Value: {}", value)
    time.Sleep( time.Second * 10 ) // wait 10 seconds!
    values <- value  // send our value back through the wormhole
}

func main() {
    fmt.Println("Go Channel Tutorial")

    // create an empty channel called "values" of type 'int'
    values := make(chan int)
    defer close(values)

    go CalculateValue(values) // pass our channel to our function

    // receiving channels will wait until a value is returned
    value := <-values // after 10 seconds a value will return
    fmt.Println(value)
}

