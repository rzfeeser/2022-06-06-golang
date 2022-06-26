/* Alta3 Research | RZFeeser
   Consuming RESTful APIs */
   
package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  // define our URL (API) as a string
  url := "https://api.spacexdata.com/v3/capsules"
  // the HTTP method we wish to send (HTTP GET)
  method := "GET"

  // For control over HTTP client headers,
  // redirect policy, and other settings,
  // create a Client
  // A Client is an HTTP client
  client := &http.Client {
  }
  
  // build a new HTTP request
  req, err := http.NewRequest(method, url, nil)

  // error checking
  if err != nil {
    fmt.Println(err)
    return
  }
  
  // if the request requires basic AUTH
  // (a protected resource), use the following:
  //                  //username  //password
  // req.SetBasicAuth("userRZF", "passQwerty!")
  
  // Send the request via a client
  // Do sends an HTTP request and
  // returns an HTTP response 
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()  // even if we have a problem, we still want this to happen

  //
  body, err := ioutil.ReadAll(res.Body)

  // error checking
  if err != nil {
    fmt.Println(err)
    return
  }
  
  // 
  fmt.Println(string(body))
}

