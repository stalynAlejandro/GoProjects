/*
STRUCTS
A struct is a collection of fields.

Structs fields are accessed using a dot.

POINTERS TO STRUCTS
Structs fields can be accessed through a struct pointer.

To access the field X of a struct whe we have the struct
pointer p we could write (*p).X

However that notation is cumbersome, so the language permit
us intead to write just p.X wihtout the explicit deference.

STRUCT LITERALS
A struct literal denotes a newly allocated struct value by
listing the values of its fields.

You can list just a subset of fileds by using the 'Name:' syntax.
And the order of named field is irrelevant.

The special prefix & returns a pointer to the struct value.

*/

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  //has type Vertex
	v2 = Vertex{X: 1}  //Y:0 is implicit
	v3 = Vertex{}      //X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// STRUCT TO POINTERS
	vv := Vertex{1, 2}
	p := &vv
	p.X = 1e9
	fmt.Println(vv)

	// STRUCT LITERALS
	fmt.Println(v1, p, v2, v3)
}
