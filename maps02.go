/* Alta3 Research | RZFeeser
   Map - hosts to IP:port resolution  */

package main

import (  
    "fmt"
)

func main() {  
    
    // create a map type, "hostResolution"
    hostResolution := map[string]string{
        "prom": "10.0.22.32:9090",
        "kafka": "10.7.8.99:9092",
        "minecraft": "192.168.59.11:25565",
    }
    
    // name to lookup
    hostName := "prom"
    
    value, ok := hostResolution[hostName]
    if ok == true {
        fmt.Println("The value of ", hostName, "is", value)
        return
    }
    fmt.Println("Socket information for", hostName, "not found")

}

