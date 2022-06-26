/* Alta3 Research | RZFeeser
   init - order of initialization  */

package main 

import "fmt"

// this will run second...
func init() {
    WhatIsThe = 0 // now it is set to 0
}


// This will run first...
var WhatIsThe = AnswerToLife()
//WhatIsThe := AnswerToLife()

func AnswerToLife() int {
    return 42
}



func main() {
    if WhatIsThe == 0 {
        fmt.Println("The cake is a lie.")
    }
}

