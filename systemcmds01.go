/* Alta3 Research | RZFeeser
   Executing system commands  */

package main

import (
    "log"
    "os/exec"
)

func main() {

    // prepares a "cmd" struct
    cmd := exec.Command("touch", "it.worked")

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}

