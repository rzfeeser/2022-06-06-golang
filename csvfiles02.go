/* RZFeeser | Alta3 Research
   Reading CSV file - ReadAll() */

package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

type User struct {
    firstName  string
    lastName   string
    occupation string
}

func main() {

    records, err := readData("users.csv")

    if err != nil {
        log.Fatal(err)
    }

    for _, record := range records {

        user := User{
            firstName:  record[0],
            lastName:   record[1],
            occupation: record[2],
        }

        fmt.Printf("%s %s is a %s\n", user.firstName, user.lastName,
            user.occupation)
    }
}

func readData(fileName string) ([][]string, error) {

    f, err := os.Open(fileName)

    if err != nil {
        return [][]string{}, err
    }

    defer f.Close()

    r := csv.NewReader(f)

    // Here we skip the first line, which contains the column names
    if _, err := r.Read(); err != nil {
        return [][]string{}, err
    }

    // get all of the records in a single swipe
    records, err := r.ReadAll()

    if err != nil {
        return [][]string{}, err
    }

    return records, nil
}

