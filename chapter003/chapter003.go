package main

import (
	"fmt"
)

//package scope
// type declaration
type List []int

// methods defined on the type
func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

//  primitive types
var y int = 42
var z string = "Let's do it!"
var a string = `She said "I don't know yaar!!"` // back quotes gives your raw string literals for mulit line comments, double quotes gives you string, strings are immutable. number of bytes is the length of the string,
// len of a string can be discovered using the len function. len is a compile time constant if the string is a constant.
var flags bool = false // boolean types us defined using bool

// Composite types
// Array types - An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.
//The length of array a can be discovered using the built-in function len. The elements can be addressed by integer indices 0 through len(a)-1. Array types are always one-dimensional but may be composed to form multi-dimensional types.

var arr1 [5]int

//Slice types
//An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.The type []T is a slice with elements of type T.
//A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low: high] -- selects a half-open range which includes the first element but excludes the last one.

func main() {

	fmt.Println("Start of Types in GO")
	/*
		The language predeclares certain type names(boolean,string,int,float). Others are introduced with type declarations.
		Composite types—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.
		Each type T has an underlying type to which it refers
	*/

	/* Method Sets: ToDo:(https://github.com/golang/go/wiki/MethodSets)
	   The method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type.
	   The method set of an interface type is its interface.
	   The method set of any other named type T consists of all methods with receiver type T.
	   The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T (that is, it also contains the method set of T).
	   Any other type has an empty method set. In a method set, each method must have a unique name.

	   Calls:
	   A method call x.m() is valid if the method set of (the type of) x contains m and the argument list can be assigned to the parameter list of m.
	   If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m()
	*/
	// A bare value
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len())

	// A pointer value
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)\n", plst, plst.Len())

	/*Boolean Types
	  A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is bool; it is a defined type.
	*/
	/*Numeric Types
	  A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are:
	    uint8       the set of all unsigned  8-bit integers (0 to 255)
		uint16      the set of all unsigned 16-bit integers (0 to 65535)
		uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
		uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

		int8        the set of all signed  8-bit integers (-128 to 127)
		int16       the set of all signed 16-bit integers (-32768 to 32767)
		int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
		int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		float32     the set of all IEEE-754 32-bit floating-point numbers
		float64     the set of all IEEE-754 64-bit floating-point numbers

		complex64   the set of all complex numbers with float32 real and imaginary parts
		complex128  the set of all complex numbers with float64 real and imaginary parts

		byte        alias for uint8
		rune        alias for int32
	*/
	//Print

	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", z)
	fmt.Println(z)
	fmt.Printf("%T \n", a)
	fmt.Println(a)
	fmt.Printf("%T \n", flags)
	fmt.Println(flags)
	fmt.Printf("%T \n", arr1)
	fmt.Println(arr1)
	arr1[4] = 100
	fmt.Println("get: ", arr1[4])
	fmt.Println("len: ", len(arr1))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("decl: ", b)

	// Two dimensional array
	var twoD [2][3]int
	// Assigning value to the array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Slices
	primes := [5]int{2, 3, 5, 7, 11}

	var s []int = primes[2:3] // includes the first element and not the last element, slices are one dimensional,
	//distict slices may share the same storage, as the underlying array is the same
	var s1 []int = primes[1:3]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println("Length of slice: ", len(s1)) // len() function in go takes in the primitive and composite types as arguments

	//A new, initialized slice value for a given element type T is made using the built-in function make,
	//which takes a slice type and parameters specifying the length and optionally the capacity.
	//A slice created with make always allocates a new, hidden array to which the returned slice value refers.
	var s3 = make([]string, 3, 5) // make is used to create an initialized slice value for a given element type T using the built-in function make.
	//
	fmt.Println(len(s3))
	fmt.Println(s3)
	fmt.Println(cap(s3)) // cap() is used to find the capacity of the slice

}
