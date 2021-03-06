/*
FOR
Go has only one looping construct, the 'for' loop.

The basic 'for' loop has three components separated by semicolons:

	The 'init' statement: executed before the first iteration.

	The 'condition' expression: evaluated before every iteration.

	The 'post' statement: executed at the end of every iteration.

The init statement will often be a short variable declaration, and the
variables declared there are visible only in the scope of the 'for' statement.

The loop will stop iterating once the boolean condition evaluates to 'false'.

FOR CONTINUED

The init and post statements are optional

IF


*/

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// If statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// IF stament can start with a short statement to execute before the
	// condition. Variables delcared by the statement to execute are
	// only in scope until the end of the if.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// The 'init' and 'post' statement are optional
	// Its the same as While
	count := 1
	for count < 1000 {
		count += count
	}
	fmt.Println(count)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}
