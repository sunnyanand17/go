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

}

func foo() {
	fmt.Println("This is the call to the foo function")
}
