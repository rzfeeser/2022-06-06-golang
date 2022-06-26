/* Alta3 Research | RZFeeser
   Interface - Abstract class */
   
package main
   
import (
    "fmt"
)

// define the interface
// this is a contract on what an animal must contain
// the goal here is to ensure we have common "tools" on anything "Animal"
type Animal interface{
    // a specialized type embedded with
    // the abstract class will define this
    Sound() string
    // the abstract class defines how to do this
    MakeSound()
}

// we never want to be able to create an abstract class
// as the goal of an abstract class is to never be instantiaed, and provide a partial
// definition of some methods (receivers)
// here abstractAnimal is working as a base-class for other specialized
// classes to embed within themselves (go calls inheritance embedding)
type abstractAnimal struct {Animal}


// this receiver (method) ensures all abstractAnimals
// will MakeSound() using the following pattern
// however, MakeSound() is only ONE of the TWO required receivers (methods)
func (a abstractAnimal) MakeSound() {
    fmt.Printf("%v", a.Sound())
}

// a specialized type called a chicken is embedded with the subclass
type Chicken struct {abstractAnimal}
func (c Chicken) Sound() string {
    return "bwaaak bwak bwak bwak\n"
}

// make it easy to make a chicken
// if you don't describe the interface as belonging to a chicken
// it will error
func NewChicken() *Chicken {
    chicken := Chicken{abstractAnimal{}}
    chicken.abstractAnimal.Animal = chicken // comment this out to error!
    return &chicken
}

type Lion struct {abstractAnimal}
func (d Lion) Sound() string {
    return "RAAAAWWWWWRRRR\n"
}

func NewLion() *Lion{
    lion := Lion{abstractAnimal{}}
    lion.abstractAnimal.Animal = lion
    return &lion
}

func main(){
    c := NewChicken()
    c.MakeSound()
    d := NewLion()
    d.MakeSound()
}  

