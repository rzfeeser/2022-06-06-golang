/* Alta3 Research | RZFeeser
   SOLUTION 02 - struct Virtmach and two methods
   ipset(string) - string is the IP to set on the struct 
   diskexpand(int) - int is the number of GB by which to expand the disk
   display() - display information about the struct      */

package main
 
import (
    "fmt"
)

type Virtmach struct{
    ip string
    hostname string
    diskgb int
    ram int
}

// ipset(string) - change IP address
// use a pointer here as we are making a change
func (v *Virtmach)ipset(ip string){
    v.ip = ip     // change the value of ip associated with Virtmach
}

// diskexpand(int) - expand disk by int
// use a pointer here as we are making a change
func (v *Virtmach)diskexpand(gb int){
    v.diskgb = v.diskgb + gb     // increase 
}

func (v Virtmach)display(){
    fmt.Println("Primary IP Address:", v.ip)    // primary IP address 
    fmt.Println("Hostname:", v.hostname)        // hostname 
    fmt.Println("Disk GB:", v.diskgb)           // diskgb 
    fmt.Println("RAM:", v.ram)                  // ram 
}
 
func main() {

    // create a virtual machine
    vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8}  // ip, hostname, diskgb, ram
    
    vm1.display()                   // show the current state
    
    vm1.diskexpand(3)               // increase by 3 GB
    
    vm1.ipset("192.168.77.33")      // change the IP address
    
    vm1.display()                   // show the changes
}

