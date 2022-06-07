/* RZFeeser | Alta3 Research
   Create a struct named person */

package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30}) // fields can be referenced by name

    fmt.Println(person{name: "Fred"})   // omitted fields are zero filled

    s := person{age: 50, name: "Okon"}  // passing by name, not by position
    fmt.Println(s.name)                 // "Okon"
    fmt.Println(s.age)                  // 50

    // & operator, or "address of operator"
    sp := &s                            // "sp" is a pointer to the actual struct "s"
    fmt.Println(sp.age)                 // 50

    // * operator for "pointer indirection"
    // For a pointer x, the pointer indirection *x denotes the value which x points to
    // Pointer indirection is rarely used, since Go can automatically take the address of a variable.
    sp.age = 51     // same as: (*sp).age = 51
    fmt.Println(sp.age)                 // pointer now reflects 51
    fmt.Println(s.age)                  // the struct also reflects 51
}

