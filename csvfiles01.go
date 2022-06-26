/* RZFeeser | Alta3 Research
   Reading CSV file to stdout - Read()  */

package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {

    // open "vendorData.csv"
    f, err := os.Open("vendorData.csv")

    if err != nil {

        log.Fatal(err)
    }

    r := csv.NewReader(f)

    for {

        // exploring the "Read()" function
        record, err := r.Read()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }

        for value := range record {
            fmt.Printf("%s\n", record[value])
        }
    }
}

