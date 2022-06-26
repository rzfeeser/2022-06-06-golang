/* Alta3 Research | RZFeeser
   Channels - Using select */
   
package main

import (
    "fmt"
    "time"
)

func main() {

    // select across two channels
    c1 := make(chan string)
    c2 := make(chan string)

    // each channel will receive a value after some amount
    // of time
    // this passage of time represents a RPC executing in
    // concurrent go routines
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // loops twice as we have two go routines
    // after both finish, the code will end
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}

