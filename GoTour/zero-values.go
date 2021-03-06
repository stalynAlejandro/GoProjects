/*
Zero Values
Variables declared without an explicit initial
value are given their zero value.

0 for númeric types,
false for the boolean type,
"" for string
*/
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q", i, f, b, s)
}
