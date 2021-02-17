/*
PACKAGES AND IMPORTS.

Every Go programs is made up of packages.
Programs start running in package main.

By convention, the package name is the same
as the last element of the import path.

For instance, the "math/rand" package comprises
files that begin wiht the statement package rand.

EXPORTED NAMES

In Go, a name is exported if it begins with a
capital letter. For example, 'Pizza' is an exported
name, as is 'Pi', which is exported from the math
package.

'pizza' and 'pi' do not start wiht capital letter
so they are not exported.

When importing a package, you can refer only to
its exported names. Any "unexported" names are not
accessible from outside the package.

*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems", math.Sqrt(7))

	//Exported names
	fmt.Printf("\nNumber pi from math is %g ", math.Pi)
}
