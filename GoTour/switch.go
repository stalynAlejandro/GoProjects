/*
SWITCH
A switch statement is a shorter way to write a sequence of if-else

Go only runs the selected case, not all cases that follow.
In effect, the 'break' statement that is needed at the end of each case
in those languages is provided automatically in Go.

Another important difference is that Go's switch cases need not be
constants, and the values involved need not be integers.

SWITCH EVALUATION ORDER

Switch cases evaluate cases from top to bottom, stopping when a case
succeeds.

swtich i {
case 0:
case f():

}

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
