// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// 1.1, 1.2
func main() {
	var input, sep string
	// same as range os.Args[1:] if len is ommited from  [1:num] it will either default to 0 or string length
	for i, char := range os.Args[1:len(os.Args)] {
		// os.Args[0 +1 onwards]
		fmt.Printf("%d, \t %s\n", i+1, char)
		input += sep + char
		sep = " "
	}
	// args 0 name of the program after 0 > user arguments
	fmt.Printf("%s: %s\n", os.Args[0], input)
}
