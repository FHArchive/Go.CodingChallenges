/*
Date 09/01/2020
Author FredHappyface
*/
package main

import (
	"testing"
)

/*
Test the swap function with "hello" and "world"
*/
func TestSwap(t *testing.T) {
	a, b := swap("hello", "world")
	if a != "world" && b != "hello" {
		t.Errorf("swap(\"hello\", \"world\") failed. Expected a = \"world\", b = \"hello\", got a = %v, b = %v", a, b)
	}
}
