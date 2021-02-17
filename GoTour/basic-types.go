/*
BASIC TYPES

Go's basic types are

bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte	 // alias for unit8
rune	 // alias for int32
		 // represents a Unicode point

float32 float64
complex64 complex128

The examples shows variables of several types,
and also that variables declarations may be "factored"
into blocks, as with import statements.

The int, uint and uintptr types are usually 32 bits
wide on 32-bit systems and 64 bits wide on 64-bit systems.

When you need an integer value you should use int unless
you have a specific reason to use a sized or unsigned integer
type.

ZERO VALUES.

Variables declared without an explicit initial value are
given their zero value.

The zero value is:
	0 		for numeric types
	false 	for the boolean type
	"" 		(the empty string) for strings


TYPE CONVERSIONS
The expression T(v) converts the value v to the type T.

Some numeric conversions:

	var i int 		= 42
	var f float64 	= float64(i)
	var u uint		= uint(f)

Or, put more simply:

	i := 42
	f := float64(i)
	u := uint(f)

Unlike in C, Go assigment between items of
different type requires an explicit conversion.

Try removing the float64 or uint conversions in
the example and see what happens.

TYPE INFERENCE

Whe declaring a variable without specifying an
explicit type (either by using ':=' syntax or var =
expressions syntax), the variable's type is inferred
from the value on th right hand side.

When the right hand side of the declaration is typed,
the new variable is the same type:

	var i int
	j := i		// j is an int

But when the right hand side contains an untyped numeric
constant, the new variable may be an int, float64 or
complex128 depending on the precision of the constant.

	i := 42				// int
	f := 3.142			// float64
	g := 0.867 + 0.5i 	// complex128

Try changing the initial value of v in the example code
and observe how its type is affected.
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// %T => type, %v => value
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)

	// Type conversions
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z)

	// Type inference
	v := 42.5 // change me
	fmt.Printf("v is type %T\n", v)
}