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
Define the swap function. Unlike some other popular languages
such as Java, its possible to have multiple outputs
this accepts two strings x and y and returns them in thr form
y, x
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
The main method is used to test the swap function
hello world is swapped to produce world hello
*/
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
