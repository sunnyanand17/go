package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// checking for errors at each step and using panic function to trigger when we can't handle an error gracefully
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fmt.Println("Start of Chapter 004")
	//Printing to standard output
	fmt.Println("We are printing to the standard output")
	//Printing to a variable
	var result string = fmt.Sprintln("we are printing to a variable")
	fmt.Println(result)
	//Reading the whole file
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat)) // Print thw whole contents of the file
	// Open a file and get the file descriptor
	f, err1 := os.Open("/tmp/dat")
	check(err1)
	// read the file in small proportions using a byte buffer
	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // read using the file descriptor and put it in the buffer and count the number of bytes read
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1)) // read the number of bytes and convert it to a string and print it.
	// Writing to a file 

}
