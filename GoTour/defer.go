/*
DEFER
A 'defer' statement defers the execution of a function until the
sorrounding function returns.

The deferred call's arguments are evaluated immediately, but the function
call is not executed until the surroundings function returns.

STACKING DEFERS
Deferred function calls are pushed onto a stack. When a function returns,
its deferred calls are executed in last-in-first-out order.

*/

package main

import "fmt"

func main() {
	defer fmt.Println("Hello World")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
