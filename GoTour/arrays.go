/*
ARRAY
The type [n]T is an array of n values of type T.

The expression

	var a[10]int

Declares a variable a as an array of then integers.

An array's length is a part of its type, so array cannot be resized.

SLICES
An array has a fixed size. A slice, on the other hand, is a dynamically-sized
flexible view into the elements of an array.

In practice, slices are much more common than arrays.

The type []T is a slice with the elements of T.

A slice is formed by specifying two indies, a low and high bound, separated by
a colon:
	a[low : high]

This selects a half-open range which includes the first element, but excludes
the las one.

The following expression creates a slice which includes elements 1 through 3 of a:

	a [1 : 4]

SLICES ARE LIKE REFERENCES TO ARRAYS

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its
underlying array.

Other slices that share the same underlying array will se those changes.

SLICES LITERALS
A slice is like an array literal without the length.

This is an array literal:

	[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it.

	[]bool{true, true, false}

SLICE DEFAULTS
Whe slicing, you may omit the high or low bounds to use their defaults instead.

The defaults is zero for the low bound and the length of the slice for the hihg bound.

For the array.

	var a [10]int

these slice expressions are equivalent:

	a[0:10]
	a[:10]
	a[0:]
	a[:]


SLICE LENGTH AND CAPACITY

A slice has both a lenght and a capacity
The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting
from the first element in the slice.

The length and capacity of a slice 'q' can be obtained using the expressions 'len(q)'
and 'cap(q)'

You can extend a slice's length by re-slicing it, provided it has sufficient capcaty.
Try changing one of the slice operations in the example program to extend it beyond
its capacity and see what happens.

NIL SLICES
The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying array.

*/
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slices
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// slices like references
	c := names[0:2]
	d := names[1:3]

	fmt.Println(c, d)

	c[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(t)

	// Slice and length and capacity
	w := []int{2, 3, 5, 7, 11, 13}
	printSlice(w)

	// Slice the slice to give it zero length
	w = w[:0]
	printSlice(w)

	// Extend its length
	w = w[:4]
	printSlice(w)

	// Drop its first two values
	w = w[2:]
	printSlice(w)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
