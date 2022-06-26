/* Alta3 Research | RZFeeser
   mySQL - Prepared statement operations */   

package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

type City struct {
    Id         int
    Name       string
    Population int
}

func main() {

    db, err := sql.Open("mysql", "student:playstation5@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    // hardcoded the "myid" integer value
    var myid int = 1

    // The ? is a placeholder which is filled with the value from the myid 
    // Under the hood, db.Query prepares, executes, and closes a prepared statement
    res, err := db.Query("SELECT * FROM cities WHERE id = ?", myid)
    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    if res.Next() {

        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%v\n", city)
    } else {

        fmt.Println("No city found")
    }
}

