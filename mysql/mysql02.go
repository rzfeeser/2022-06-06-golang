/* Alta3 Research | RZFeeser
   mySQL - Select all rows and display data */

package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// build a struct we can use to extract data from our
// mySQL instance
type City struct {
    Id         int
    Name       string
    Population int
}

func main() {

    // connect to mySQL
    db, err := sql.Open("mysql", "student:playstation5@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    // here is our magic
    res, err := db.Query("SELECT * FROM cities")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    // Next prepares the next result row for reading with the Scan method
    // returns true on success
    // returns false if there is no next result row 
    // returns error if happens during prep
    for res.Next() {

        // create a struct called, city
        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%v\n", city)  // %v is the value in a default format
    }
}

