// Exercise description
// Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64.
// Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values.

package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring and assigning the max legal value
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI int64 = math.MaxInt64
	// Incrementing variables.
	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)

}
