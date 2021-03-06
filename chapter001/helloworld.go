package main

import (
	"fmt"
)

// entry point to a go program
func main() {
	fmt.Println("Hello World")

	// Control flow sequence
	foo()
	// Control flow loop;iterattive

	for i := 0; i < 10; i++ {
		//Control flow conditional

		if i%2 == 0 {
			fmt.Printf("This is an even number %v \n", i)
		}
	}
	// Print and Scan functions of the fmt package. Lets review them
	/* Documentation Link:https://golang.org/pkg/fmt/#Errorf
	//Sample examples below
	//Genreal Verbs
	%v	the value in a default format ,when printing structs, the plus flag (%+v) adds field names
	%#v	a Go-syntax representation of the value
	%T	a Go-syntax representation of the type of the value
	%%	a literal percent sign; consumes no value
	Default format for %v translates to below verbs for the different types
	bool:                    %t
	int, int8 etc.:          %d
	uint, uint8 etc.:        %d, %#x if printed with %#v
	float32, complex64, etc: %g
	string:                  %s
	chan:                    %p
	pointer:                 %p


	// Boolean
	%t	the word true or false

	//integers
	%b	base 2
	%c	the character represented by the corresponding Unicode code point
	%d	base 10
	%o	base 8
	%q	a single-quoted character literal safely escaped with Go syntax.
	%x	base 16, with lower-case letters for a-f
	%X	base 16, with upper-case letters for A-F
	%U	Unicode format: U+1234; same as "U+%04X"

	//Floating point and complex constituents
	%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
	%e	scientific notation, e.g. -1.234456e+78
	%E	scientific notation, e.g. -1.234456E+78
	%f	decimal point but no exponent, e.g. 123.456
	%F	synonym for %f
	%g	%e for large exponents, %f otherwise. Precision is discussed below.
	%G	%E for large exponents, %F otherwise

	//String and slice of bytes (treated equivalently with these verbs)
	%s	the uninterpreted bytes of the string or slice
	%q	a double-quoted string safely escaped with Go syntax
	%x	base 16, lower-case, two characters per byte
	%X	base 16, upper-case, two characters per byte

	//Slice
	%p	address of 0th element in base 16 notation, with leading 0x

	//Pointers
	%p	base 16 notation, with leading 0x
	The %b, %d, %o, %x and %X verbs also work with pointers,
	formatting the value exactly as if it were an integer.

	//Compound objects - the elements are printed recursivley using these below rules
	struct:             {field0 field1 ...}
	array, slice:       [elem0 elem1 ...]
	maps:               map[key1:value1 key2:value2 ...]
	pointer to above:   &{}, &[], &map[]

	// Explicit argument indexes:
	the notation [n] immediately before the verb indicates that the nth one-indexed argument is to be formatted instead

	//All errors begin with the string "%!" followed sometimes by a single character (the verb) and end with a parenthesized description.

	//blank identifier is represented by the underscore in go. It can be used to assign values for operands which will be discarded by go harmlessly.

	*/
	s := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	fmt.Print(s)
}

func foo() {
	fmt.Println("This is the call to the foo function")
}
