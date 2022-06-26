/* Alta3 Research | RZFeeser
   Channels - Intro to buffered channels*/

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
    // after the data is through the wormhole
    fmt.Println("I run after the data is added to the channel")
}

func main() {
    fmt.Println("Go Channel Tutorial")

    // create an empty channel called "values" with a buffer of 4
    values := make(chan int, 4)
    defer close(values)

    go CalculateValue(values) // pass our channel to our function
    go CalculateValue(values)
    go CalculateValue(values)
    go CalculateValue(values)

    // wait until a value is returned
    value := <-values // after 10 seconds a value will return

    time.Sleep( time.Second ) // wait a single second (after data is read from the channel)
    fmt.Println(value) // I run after the data is read from the channel
}

