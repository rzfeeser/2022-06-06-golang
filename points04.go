/* RZFeeser | Alta3 Research
   Without Pointers - Why we need pointers  */    

package main

import (
 "fmt"
)

type User struct {
 Name string
 Pets []string
}

func (u User) newPet() {
 u.Pets = append(u.Pets, "Lucy")        // Simple function should ensure "Lucy" is added to User
 fmt.Println(u)                         // This user is **NOT** the same user as the one in main()!!
}

func main() {
 u := User{Name: "Anna", Pets: []string{"Bailey"}}
 u.newPet()                                         // {Anna [Bailey Lucy]} -- This *should* add "Lucy" to "u"
 fmt.Println(u)                                     // We **DO NOT** see Lucy -- what happened?
}

