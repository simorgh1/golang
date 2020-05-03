package main

import "fmt"

func main() {

	str := "Hello World!"

	fmt.Printf("Invert the ascii string: '%s'\n", str)

	// string in go is immutable, convert it first to rune before inverting.
	buf := []rune(str)

	// traverse through the last char to the first and replace them
	for i, _ := range str {
		buf[i] = rune(str[len(str)-1-i])
	}

	// convert back to string
	str = string(buf)
	fmt.Printf("Result: '%s'\n", str)
}
