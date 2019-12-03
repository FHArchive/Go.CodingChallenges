// 2>/dev/null; $(which go) run $0 $@; exit $?
/*
Date 08/10/2018
Author FredHappyface
*/
package main

import (
	"fmt"
	"runtime"
)

/*
The main method identifies the OS that go is running
on
*/
func main() {
	fmt.Print("Go is running on ")
	os := runtime.GOOS
	// Here is an example of a switch-case statement
	switch os {
	case ("darwin"):
		{
			fmt.Print("OSX")
		}
	case ("linux"):
		{
			fmt.Print("Linux")
		}
	default:
		{
			fmt.Printf("%s", os)
		}
	}
}
