/*
VARIABLES

The var statement declares a list of varibales.
As in function argument lists, the type is last.

A var statement can be package or function level.
We see both in this example.

VARIABLES WITH INITIALIZERS

A var declaration can include initializers,
one per variable.

If an initializer is present, the type can be omitted;
The variable will take the type of the initializer.

SORT VARIABLE DECLARATIONS

Inside a function, the ':=' short asigment statement can
be used in place of a 'var' declaration with implicit type.

Outside a function, every statement begins with a keyword
('var', 'func' and so on) and so the ':=' construct is
not available.
*/

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// Short declarations
	k := 3
	fmt.Println(k)
}
