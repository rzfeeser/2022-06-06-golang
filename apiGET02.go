/* Alta3 Research | RZFeeser
   API - RESTful API lookup, and decoding JSON on 200 response  */
package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
        "time"
)


type SpaceX []struct {
        CapsuleSerial      string    `json:"capsule_serial"`
        CapsuleID          string    `json:"capsule_id"`
        Status             string    `json:"status"`
        OriginalLaunch     time.Time `json:"original_launch"`
        OriginalLaunchUnix int       `json:"original_launch_unix"`
        Missions           []struct {
                Name   string `json:"name"`
                Flight int    `json:"flight"`
        } `json:"missions"`
        Landings   int    `json:"landings"`
        Type       string `json:"type"`
        Details    string `json:"details"`
        ReuseCount int    `json:"reuse_count"`
}

func main() {
        // define our URL (API) as a string
        url := "https://api.spacexdata.com/v3/capsules"

        // Build the request
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                log.Fatal("NewRequest: ", err)
                return
        }

        // For control over HTTP client headers,
        // redirect policy, and other settings,
        // create a Client
        // A Client is an HTTP client - you can load it up with
        // parameters if you wanted (timeouts, tls, retries, etc.)
        // You only need ONE of these (supported by goroutines)
        client := &http.Client{}

        // Send the request via a client
        // Do sends an HTTP request and
        // returns an HTTP response
        resp, err := client.Do(req)
        if err != nil {
                log.Fatal("Do: ", err)
                return
        }

        // Callers should close resp.Body
        // when done reading from it
        // Defer the closing of the body
        defer resp.Body.Close()

        // Fill the record with the data from the JSON
        var record SpaceX

        // Use json.Decode for reading streams of JSON data
        if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
                log.Println(err)
        }

        // the _ is an int
        // launchData is the data
        for launchNo, launchData := range record {
        fmt.Println("Capsule Record     =\n", launchData)
        fmt.Println("Record Number      =", launchNo)
        fmt.Println("Capsule Serial     =", launchData.CapsuleSerial)
        }
}

