/* RZFeeser | Alta3 Research
  Getting at interfaces */ 

package main

import (
    "fmt"
    "math"
)

// basic interface for geometric shapes
type geometry interface {
    area() float64              // needs an area method that returns a float64
    perim() float64             // needs a perim method that returns a float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

// To implement an interface in Go we just need to implement all the methods in the interface
// Here we implement geometry on rects.
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// the implementation for circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// a generic "measure" function taking advantage of the interface "geometry"
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)    // run measure against rect
    measure(c)    // run measure against circle
}

