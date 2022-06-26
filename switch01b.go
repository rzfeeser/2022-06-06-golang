/* RZFeeser | Alta3 Research
   CHALLENGE 01 - match on any minor version of Go  */

package main
  
import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

    // init gove
    var gove string = runtime.Version()  // this returns the version of Go
    
    switch {
    case strings.Contains(gove, "go1.18"):      // if matches "go1.18", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case strings.Contains(gove, "go1.17"), strings.Contains(gove, "go1.16"):       // if matches "go1.17" OR "go1.16"
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("Upgrade GoLang before you continue")
    }
}

