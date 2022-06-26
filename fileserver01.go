/* RZFeeser | Alta3 Research
   HTTP File server            */
 
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {

    // A file server is registered with Handle() and serves files from the public/ directory
    fileServer := http.FileServer(http.Dir("./public"))
    http.Handle("/", fileServer)

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello there!\n")
    })

    fmt.Printf("Starting server at port 8044\n")
    log.Fatal(http.ListenAndServe(":8044", nil))
}     

