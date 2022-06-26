/* Alta3 Research | RZFeeser
   Mixing tests with our benchmark tests. */
   
package main

import (
    "fmt"
    "testing"
)

// additional test (non benchmark)
func TestCalculate(t *testing.T) {
    fmt.Println("Test Calculate")
    expected := 4
    result := Calculate(2)
    if expected != result {
        t.Error("Failed")
    }
}

// additional test (non benchmark)
func TestOther(t *testing.T) {
    fmt.Println("Testing something else")
    fmt.Println("This shouldn't run with -run=calc")
}

// called by other 3 functions
func benchmarkCalculate(input int, b *testing.B) {
    for n := 0; n < b.N; n++ {
        Calculate(input)
    }
}

// function calls with different types of input
func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }            

