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

	// Topics -- Tokens, Comments, Identifier, Keywords,Operators & Punctuations

}
