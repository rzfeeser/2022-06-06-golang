/* Alta3 Research | RZFeeser
   goroutine - go routines without waitgroups   */

package main

import (
        "fmt"
        "time"
)

func countDown(s int) {
        for i := s; i > 0; i-- {
                fmt.Println(i)
                time.Sleep(time.Second)   // wait one second
        }
}


func main() {

    fmt.Println("NASA launch sequence starting:")

    // Begin the countdown
    go countDown(10)

    fmt.Println("Launch sequence starting:")

    time.Sleep(time.Second)
    fmt.Println("Hey wait a second...")
    
    time.Sleep(time.Second)
    fmt.Println("I forgot my wallet!")


    fmt.Println("TAKE OFF!")
}

