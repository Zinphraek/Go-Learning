package main

import "fmt"

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "Hello"
	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x) // Will not compile
	fmt.Println(y) // Will not compile
}
