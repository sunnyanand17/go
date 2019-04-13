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
	*/

	//4. Literals

}
