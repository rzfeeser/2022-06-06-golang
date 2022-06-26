/* RZFeeser | Alta3 Research
   HTTP POST request with Form data  */

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
)

func main() {

    data := url.Values{
        "name":       {"John Doe"},
        "occupation": {"gardener"},
    }

    // Send the HTTP POST request
    resp, err := http.PostForm("http://httpbin.org/post", data)

    if err != nil {
        log.Fatal(err)
    }

    // Decode the response body into a map
    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    // print the received data
    fmt.Println(res)
}

