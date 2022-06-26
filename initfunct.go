/* Alta3 Research | RZFeeser
   init function              */

package main

import "fmt"

func init() {
  fmt.Println("Think of the init function like a prelude to main.")
}

func main() {
  fmt.Println("This would run 'after' the init.")
}

