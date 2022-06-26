/* Alta3 Research | RZFeeser
   goroutine - with waitgroups   */

package main

import (
        "fmt"
        "time"
        "sync"
)

func countDown(s int) {
        defer wg.Done()
        for i := s; i > 0; i-- {
                fmt.Println(i)
                time.Sleep(time.Second)   // wait one second
        }
}


// WaitGroup is used to wait for the program to finish goroutines
var wg sync.WaitGroup

func main() {

    fmt.Println("NASA launch sequence starting:")

    // the waitGroup has "1" in the queue
    wg.Add(1)

    // Begin the countdown
    go countDown(10)

    fmt.Println("Launch sequence starting:")

    time.Sleep(time.Second)
    fmt.Println("Hey wait a second...")
    
    time.Sleep(time.Second)
    fmt.Println("I forgot my wallet!")

    // wait for the waitGroup queue to reach 0
    wg.Wait()

    fmt.Println("TAKE OFF!")
}

