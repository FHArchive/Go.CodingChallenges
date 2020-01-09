/*
Date 09/01/2020
Author FredHappyface
*/
package main

import (
	"testing"
)

/*
Test the intToString function with 123
*/
func TestIntToString(t *testing.T) {
	result := intToString(123)
	if "one two three " != result {
		t.Errorf("intToString(123) failed. Expected \"one two three \", got \"%v\"", result)
	}
}
