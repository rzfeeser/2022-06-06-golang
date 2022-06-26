/* Alta3 Research | RZFeeser
   Channels and goroutines - skeleton for HTTP lookups with 3 workers */

package main

import (
        "fmt"
        "sync"
        "time"
)

// this is our "worker" function
// the worker wants to harvest data from the channel (linkChan)
func worker(linkChan chan string, wg *sync.WaitGroup) {
        // Decreasing internal counter for wait-group as soon as goroutine finishes
        defer wg.Done()      

        // loop across the incoming channel (it is treated as a slice)
        for url := range linkChan {
                time.Sleep(1 * time.Second)                     // this would be work we want to perform
                fmt.Printf("Done processing link #%s\n", url)   // this would be work we want to perform
        }

}


func main() {
        // imagine we have a slice of links we want to perform HTTP GET requests on
        yourLinksSlice := make([]string, 50)
        
        // write 1 to 50 in the slice we just created
        for i := 0; i < 50; i++ {
                yourLinksSlice[i] = fmt.Sprintf("%d", i+1)
        }

        // use make or the channel won't be useful
        lCh := make(chan string)
        
        
        wg := new(sync.WaitGroup) // returns a pointer to initialized memory

        // Adding routines to workgroup and running then
        for i := 0; i < 3; i++ {
                wg.Add(1)
                go worker(lCh, wg)
        }

        // Processing all links by spreading them to `free` goroutines
        for _, link := range yourLinksSlice {
                lCh <- link
        }

        // Closing channel (waiting in goroutines won't continue any more)
        close(lCh) // no more jobs will be sent here

        // Waiting for all goroutines to finish (otherwise they die as main routine dies)
        wg.Wait()
}

