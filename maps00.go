/* Alta3 Research | RZFeeser
   Maps - Review of Maps using static types */

package main

import (
    "fmt"
)

func main() {
    
    // make assigns a memory address (will not return nil)
    var gamescores = make(map[string]int)
    gamescores["Zerg"] = 9092

    // Returns gamescores: map[Zerg:9092]
    fmt.Println("gamescrores:", gamescores)

    //var totalscore map[string]int   // i will return Nil, comment out to make work

    // Runtime error: panic: assignment to entry in nil map
    //totalscore["minecraft"] = 912   // comment out this line
    
    // if I want to make a map without filling with data right away
    // use make()
    var network = make(map[string]int)
    network["Cisco"] = 190
    network["Arista"] = 56
    network["NetGear"] = 302

    // Returns network: map[Cisco:190 Arista:56 NetGear:302]
    fmt.Println("Network Devices:", network)
    
    values := map[string]int{
        "abc": 123,
        "def": 345,
        "ghi": 567,
        "jkl": 897,
        "mno": "larry",
    }
    // Returns values: map[abc:123 def:345 ghi:567 jkl:897]
    fmt.Println("values:", values)
    
}

