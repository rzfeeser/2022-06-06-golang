/* Alta3 Research | RZFeeser
   Testing - a simple test file */

package main

import (
    "testing"   // used for testing
    "regexp"    // regular expression library
)

func TestHello(t *testing.T) {
    testName := "RZFeeser"
    desiredResult := regexp.MustCompile("Hello there "+testName)

    result, err := Hello("RZFeeser")
    
    if !desiredResult.MatchString(result) || err != nil {
        t.Fatalf("Wanted %v, but got, %v, or got %v", desiredResult, result, err)
    }
}

