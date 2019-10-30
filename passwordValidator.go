/*
By Kieran BW
Psk is valid if MinLength = 8, MaxLength = 64, 1 Int, 1 Special, 0 Spaces
*/

package main

import (
	"fmt"
	"os"
)

/*
main iterates through a list of passwords and for each, prints the password
and if it is valid
*/
func main() {
	args := os.Args[1:]
	for _, arg := range args {
		typeMap := [3]int{0, 0, 1}
		for _, char := range arg {
			typeMap = updateTypeMap(int(char), typeMap)
		}
		valid := "False"
		if isValid(len(arg), typeMap) {
			valid = "True"
		}
		fmt.Printf("Password: %s, Valid? %s\n", arg, valid)
	}
}

/*
updateTypeMap increments special chars and numbers and decrements spaces
if these are present (spaces is decremented as these are unwanted and so
index < 1 would be satisfied if spaces are present (not valid ))
*/
func updateTypeMap(char int, typeMap [3]int) [3]int {
	if char > 47 && char < 58 {
		typeMap[0]++
	}

	if char > 32 && char < 48 {
		typeMap[1]++
	}

	if char == 32 {
		typeMap[2]--
	}

	return typeMap
}

/*
isValid outputs a bool based on password attributes.
*/
func isValid(length int, typeMap [3]int) bool {
	for _, index := range typeMap {
		if index < 1 {
			return false
		}
	}

	if length > 64 || length < 8 {
		return false
	}

	return true
}
