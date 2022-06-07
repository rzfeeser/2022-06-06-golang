/* Alta3 Research | RZFeeser
   Methods - defining bookup() to increase the value of books by 1  */

// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    books     int
    salary    int
}

// increase value of books by 1
func (a author) bookup() {
    a.books++        // increment name by 1
}
 
// Method with a receiver
// of author type
func (a author) show() {
 
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Number of books authored: ", a.books)
    fmt.Println("Salary: ", a.salary)
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:      "RZFeeser",
        branch:    "Pennsylvania",
        books:     14,
        salary:    34000,
    }
 
    // Calling the method
    res.show()
    
    // run our new method
    res.bookup()   // increase books by 1
    
    res.show()     // display the new values stored within struct
}

