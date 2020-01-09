// 2>/dev/null; $(which go) run $0 $@; exit $?
/*
Date 08/10/2018
Author FredHappyface
*/
package main

import (
	"fmt"
)

/*
Method intToString
Inputs int: input
Outputs string: output
This function takes an integer input and returns it in word
form. For example, 2809 would be "two eight zero nine"
*/
func intToString(input int) string {
	// Initialise variables (go can be weakly typed)
	output := ""
	// Initialise intAsString by converting the input to a string
	intAsString := fmt.Sprintf("%d", input)
	// Take each int and append the word version to the output
	for index := 0; index < len(intAsString); index++ {
		ascii := intAsString[index]
		switch ascii {
		case (48):
			{
				output += "zero "
			}
		case (49):
			{
				output += "one "
			}
		case (50):
			{
				output += "two "
			}
		case (51):
			{
				output += "three "
			}
		case (52):
			{
				output += "four "
			}
		case (53):
			{
				output += "five "
			}
		case (54):
			{
				output += "six "
			}
		case (55):
			{
				output += "seven "
			}
		case (56):
			{
				output += "eight "
			}
		case (57):
			{
				output += "nine "
			}
		}
	}
	// Return the output to the calling program
	return output
}

/*
	The main method is used to test the intToString function
	the program will print "one zero" with the test of 10
*/
func main() {
	output := intToString(10)
	fmt.Println(output)
}
