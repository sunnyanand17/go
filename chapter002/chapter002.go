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
	The first character in an identifier must be a letter.

	Go has 4 types of Identifiers:
	1) Blank Indentifier: is represented by the underscore character _.

	*/
	//2. Keywords

	//3. Operators & Punctuations

	//4. Literals

}
