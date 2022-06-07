/* Alta3 Research | RZFeeser
   CHALLENGE 01 - using fmt.Sprintf() to construct
   https://example.org:6001/v2/snacks?req=snickers&quantity=1&size=king
   
   `uri` is `https://example.org:6001/v2/snacks?`
   `r`   is `req=snickers`
   `q`   is `quantity=1`
   `s`   is `size=king`         */

package main

import (
    "fmt"
)

func main() {

    // declare a const uri (this will never change)
    const uri = "https://example.org:6001/v2/snacks?"

    // declare our parameters
    // this is a group declaration
    var r, q, s string = "req=snickers", "quantity=1", "size=king" 

    // create the string, but don't render it yet
    res := fmt.Sprintf("%s%s&%s&%s", uri, r, q, s)
    fmt.Println(res) // finished URI
}

