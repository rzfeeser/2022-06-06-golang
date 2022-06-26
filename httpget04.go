/* RZFeeser | Alta3 Research 
   HTTP GET and asynchronous requests */

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "regexp"
  //"sync"
)

func main() {
  
  urls := []string{
    "http://example.com",
    "https://alta3.com",
    "https://www.python.org",
  }

  // We want to make multiple asynchronous HTTP requests
  //var wg sync.WaitGroup

  for _, u := range urls {
    
  //  wg.Add(1)
    
    // A goroutine is created with the go keyword
    go func(url string) {
    
      // WaitGroups are used to manage goroutines
      // It waits for a collection of goroutines to finish
  //    defer wg.Done()
    
      content := doReq(url)
      title := getTitle(content)
      fmt.Println(title)
    }(u)
  }

  //wg.Wait()
}

func doReq(url string) (content string) {

  resp, err := http.Get(url)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatal(err)
  }

  return string(body)
}

func getTitle(content string) (title string) {

  // searching for this regex pattern (note the capturing group)
  re := regexp.MustCompile("<title>(.*)</title>")

  // perform a search on "content" (against our pattern "re")
  parts := re.FindStringSubmatch(content)

  // parts will only have len if it found matches
  // parts[0] would be the "full match"
  // parts[1] is "just" the capturing group
  // a capturing group is what is inside the parentheses of our regex expression 
  if len(parts) > 0 {
    return parts[1]
  } else {
    return "no title"
  }
}

