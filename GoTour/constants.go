/*
CONSTANTS

Constants are declared like variables, but with
the 'const' keyword.

Constant can be character, string, boolean, or
numeric values.

Constants cannot be declared using the ':=' syntax.

NUMERIC CONSTANTS

Numeric constants are high-precision values.

An untyped constant takes the type needed by its
context.

Try printing needInt(Big) too.

An int can store at maximum a 64-bit integer, and
sometimes, less.
*/

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "MANDALORIAN"

	fmt.Println("Hello", World)
	fmt.Println("HAppy ", Pi, " day")

	const Truth = true
	fmt.Println("Go, rules? ", Truth)
}
