// Declare a main package
// (a package is a way to group functions, and it's made up of all the files in the same directory).
package main

// Import the fmt package (shorthand for format). It allows us to print to the console.
import (
	"fmt"

	"rsc.io/quote"
)

// Define a main function (executes by default when we run the file)
func main() {
	fmt.Println(quote.Hello())
	Goodbye()
}

// Goodbye function
func Goodbye() {
	fmt.Println("Goodbye, World!")
}
