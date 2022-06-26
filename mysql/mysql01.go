/* Alta3 Research | RZFeeser
   mySQL - Connecting to the database  */
   
package main

import (
    "database/sql"
    "fmt"
    "log"

    // when a package is imported prefixed with a blank identifier
    // the init function of the package is called
    // the function registers the driver
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    // open the database
                        // db name,   connection information
    // db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
    db, err := sql.Open("mysql", "student:playstation5@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()    // idiomatic coding! (best practice)
                        // defer will close the sql.DB if it should
                        // not have a lifetime beyond the scope of the function 

    if err != nil {
        log.Fatal(err)
    }

    var version string

    // determine the version of mySQL
    // QueryRow executes a query that is expected to return at most one row
    // Scan function copies the column from the matched row
    // into the version variable
    err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println(version)
}

