// Exercise description:
// Write a program that declares an integer variable called i with the value 20.
// Assign i to a floating-point variable named f. Print out i and f.

package main

import "fmt"

func main() {
	i := 20
	var f float64 = float64(i)

	fmt.Println(i, f)
}
