package main

import (
	"fmt"
)

func main() {

	fmt.Print("Beginning of Chapter 2")
	/* Go Syntax
	Go uses EBNF- Extended Backnus Naur form for expressing its Context Free Grammar and language syntax.
	*/

	/* Source Code Representation
	Source code is Unicode text encoded in UTF-8. The text is not canonicalized, so a single accented code point is distinct
	from the same character constructed from combining an accent and a letter; those are treated as two code points.
	For simplicity, this document will use the unqualified term character to refer to a Unicode code point in the source text.
	Each code point is distinct; for instance, upper and lower case letters are different characters.
	*/

	//Letters and Digits in go. Blank identifier or underscore is a letter.
	/*
		letter        = unicode_letter | "_" . //Unicode code point classified as letter
		decimal_digit = "0" … "9" . // Unicode code point classified as digit
		octal_digit   = "0" … "7" .
		hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
	*/

	// Tokens

	//Comments
	/* Line Comments - starts with // and stops at the end of line
	// General Comments - start with the character sequence /* and stop with the first subsequent character sequence
	Comment cannot start inside a commnet, string or rune.
	*/

	// There are 4 types of tokens in go
	//White spaces in go are formed by the follwoing Unicode Code Points:
	// spaces, horizontal tabs, carriage returns and newline characters

	//Semicolon is the terminator for the go grammar. Go uses the below 2 rules for inserting a semi-colon.
	/*
		1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a
			line's final token if that token is:
			an identifier
			an integer, floating-point, imaginary, rune, or string literal
			one of the keywords break, continue, fallthrough, or return
			one of the operators and punctuation ++, --, ), ], or }
		2. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing ")" or "}".

	*/

	//1. Identifier
	/*Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits.
	The first character in an identifier must be a letter(unicode letter+ underscore)

	Go has 3 types of Identifiers:
	1) Blank Indentifier: is represented by the underscore character _.
	2) Predeclared Identifiers are implicitly declared in the universe block:
		Types:	bool byte complex64 complex128 error float32 float64
				int int8 int16 int32 int64 rune string
				uint uint8 uint16 uint32 uint64 uintptr

		Constants:
				true false iota

		Zero value:
				nil

		Functions:
				append cap close complex copy delete imag len
				make new panic print println real recover
	3)  Exported Identifiers: An identifier is exported to permit access to it from another package. Two rules are followed
							  while exporting identifiers:
		1) the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
		2) the identifier is declared in the package block(declared outside of any function) or it is a field name(in struct field name must start with Capital Letter)
		   or method name(in an interface or package the method name must start with a Capital Letter).


	*/
	//2. Keywords are reserved identifiers and cannot be used for program purposes other than they are defined for
	/*	break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var
	*/

	//3. Operators & Punctuations
	/*
			Operators combine operands into expressions
			Unary Operators: "+" | "-" | "!" | "^" | "*" | "&" | "<-"
			Relational operators: "==" | "!=" | "<" | "<=" | ">" | ">="
			Addition Operators: "+" | "-" | "|" | "^"
			Multiplication Operators: "*" | "/" | "%" | "<<" | ">>" | "&" | "&^"
			Binary Operators: "||" | "&&" | rel_op | add_op | mul_op
			Comparisons operators: ==    equal, !=    not equal, <     less ,<=    less or equal ,>     greater ,>=    greater or equal

			For Binary operators the operands type must be identical, unless the operation involves shifts or untyped constants.
			For operations involving constants , follow rules for constant expressions.
			Constant expressions:
								Constant expressions may contain only constant operands and are evaluated at compile time.
								Except for shift operations, if one operand is an untyped constant and the other operand is not, the constant
								 is implicitly converted to the type of the other operand.
								The right operand in a shift expression must have unsigned integer type or be an untyped constant representable
								 by a value of type uint.
								If the left operand of a non-constant shift expression is an untyped constant, it is first implicitly converted
								 to the type it would assume if the shift expression were replaced by its left operand alone.
			Arithmetic Operators in Go: Arithmetic operators apply to numeric values and yield a result of the same type as the first operand.
			+          sum                  integers, floats, complex values, strings
			-          difference           integers, floats, complex values
			*          product              integers, floats, complex values
			/          quotient             integers, floats, complex values

			%          remainder            integers
			&          bitwise AND          integers
			|          bitwise OR           integers
			^          bitwise XOR          integers
			&^         bit clear (AND NOT)  integers

			<<         left shift           integer << unsigned integer
			>>         right shift          integer >> unsigned integer



		Go Constants Overiew:
		Go is a statically typed language that does not permit operations that mix numeric types.
		You can't add a float64 to an int, or even an int32 to an int. Yet it is legal to write 1e6*time.Second or math.Exp(1) or even 1<<('\t'+2.0).
		In Go, constants, unlike variables, behave pretty much like regular numbers. There is no mixing of types in Go.
		In Go, const is a keyword introducing a name for a scalar value such as 2 or 3.14159 or "scrumptious". Constants can be built from constant expressions.
		There are many kinds of constants in Go:
		1) Integers
		2) Floats: he default type for a floating-point constant is float64, although an untyped floating-point constant can be assigned to a float32 value just fine.
		3) Runes
		4)Signed
		5)Unsigned
		6)Imaginary
		7)Complex
		8)String: Some text enclosed in double qoutes(""). Raw string literals are enclosed by backquotes (``).
		Untyped String Constant: const hello = "Hello, 世界" is an untyped string constant, that it is a constant textual value that does not have a fixed type.
		This constant hello does not have a type of string in Go and will be used as only textual value.
		An untyped constant is just a value, one not yet given a defined type that would force it to obey the strict rules that prevent combining differently typed values.
		Typed String Constant: const typedHello string = "Hello, worked!" is a typed string constant as we have explicitly defined the type this means typedHello
		constant has a Go type string and cannot explicitly be assigned any other type. Assigning untyped variable to a variable of any type compatible with strings works without error.
		It does, however, have a default type that is exposed when, and only when, no other type information is available.

		Defualt types for numeric constants:
		Integer constants default to int, floating-point constants float64, rune constants to rune (an alias for int32), and imaginary constants to complex128

		9)Boolean:The values true and false are untyped boolean constants that can be assigned to any boolean variable, but once given a type, boolean variables cannot be mixed:
		https://blog.golang.org/constants
	*/

	//4. Literals

}
