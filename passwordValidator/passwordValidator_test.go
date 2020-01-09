/*
Date 09/01/2020
Author FredHappyface
*/

package main

import (
	"testing"
)

var TypeMap = [3]int{0, 0, 1}

/*
Test the updateTypeMap function with 'a'
*/
func TestUpdateTypeMapLettera(t *testing.T) {
	result := updateTypeMap('a', TypeMap)
	expected := [3]int{0, 0, 1}

	if expected != result {
		t.Errorf("updateTypeMap('a', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with 'z'
*/
func TestUpdateTypeMapLetterz(t *testing.T) {
	result := updateTypeMap('z', TypeMap)
	expected := [3]int{0, 0, 1}

	if expected != result {
		t.Errorf("updateTypeMap('z', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with 'A'
*/
func TestUpdateTypeMapLetterA(t *testing.T) {
	result := updateTypeMap('A', TypeMap)
	expected := [3]int{0, 0, 1}

	if expected != result {
		t.Errorf("updateTypeMap('A', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with 'Z'
*/
func TestUpdateTypeMapLetterZ(t *testing.T) {
	result := updateTypeMap('Z', TypeMap)
	expected := [3]int{0, 0, 1}

	if expected != result {
		t.Errorf("updateTypeMap('Z', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with '0'
*/
func TestUpdateTypeMapNumber0(t *testing.T) {
	result := updateTypeMap('0', TypeMap)
	expected := [3]int{1, 0, 1}

	if expected != result {
		t.Errorf("updateTypeMap('0', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with '9'
*/
func TestUpdateTypeMapNumber9(t *testing.T) {
	result := updateTypeMap('9', TypeMap)
	expected := [3]int{1, 0, 1}

	if expected != result {
		t.Errorf("updateTypeMap('9', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with '!'
*/
func TestUpdateTypeMapSymbol(t *testing.T) {
	result := updateTypeMap('!', TypeMap)
	expected := [3]int{0, 1, 1}

	if expected != result {
		t.Errorf("updateTypeMap('!', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the updateTypeMap function with ' '
*/
func TestUpdateTypeMapSpace(t *testing.T) {
	result := updateTypeMap(' ', TypeMap)
	expected := [3]int{0, 0, 0}

	if expected != result {
		t.Errorf("updateTypeMap(' ', TypeMap) failed. Expected %v, got %v", expected, result)
	}
}

/*
Test the isValid function with 8
*/
func TestIsValid8(t *testing.T) {
	if isValid(8, TypeMap) {
		t.Error("TestIsValid8 failed")
	}
}

/*
Test the isValid function with 64
*/
func TestIsValid64(t *testing.T) {
	if isValid(64, TypeMap) {
		t.Error("TestIsValid64 failed")
	}
}

/*
Test the isValid function with 000
*/
func TestIsValid000(t *testing.T) {
	if isValid(10, [3]int{0, 0, 0}) {
		t.Error("TestIsValid000 failed")
	}
}

/*
Test the isValid function with 111
*/
func TestIsValid111(t *testing.T) {
	if !isValid(10, [3]int{1, 1, 1}) {
		t.Error("TestIsValid111 failed")
	}
}
