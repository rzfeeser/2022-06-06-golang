/* Alta3 Reseach | RZFeeser
   Goroutine - An examination of channels that send and channels that receive.
   chan<- is channel that is SENDING (the arrow is dumping into the wormhole)
   <-chan is a channel that is RECEIVE (the arrow flying out of the wormhole) */
   
package main

import "fmt"

// the ChanSendOnly function ONLY accepts a channel that is SENDING only (chan<- on right)
// If you try to receive on this channel, your code will not compile
func ChanSendOnly(data chan<- string, msg string){
    msg = msg + "\nvia: tag1"
    data <- msg  // i want to write this message into the channel (send it INTO the wormhole)
}

// The ChanRevOnly function can receive on the "data" channel (<-chan)
// and it can also send on the "home" channel
func ChanRecvOnly(data <-chan string, home chan<- string) {
    msg := <-data  // read from the buffer
    msg = msg + "\nvia: tag2"
    home <- msg    // send msg into the wormhole
}

func main() {
    sender := make(chan string, 1)  // create channel with buffer of one slot
    recv := make(chan string, 1)    // create channel with buffer of one slot
    
    // call the function that is ONLY set up to accept a SENDING channel
    // we know this channel is SENDING only because the arrow is to the right of chan
    ChanSendOnly(sender, "I can pass a message into this function")
    
    // at this point, the augmented message is in the channel buffer

    // call the function that can receive (arrow to the left of chan)
    // and can send (arrow to the right of chan)
    ChanRecvOnly(sender, recv)

    // Pull data out of the channel and print it to the screen
    fmt.Println(<-recv)
    // Displays:
    // I can pass a message into this function
    // via: tag1
    // via: tag2
}

