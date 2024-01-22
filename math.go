package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	// augmented assignment
	var d = 0
	d += 10 // d = d + 10
	fmt.Println(d)

	d += 5 // d = d + 5
	fmt.Println(d)

	d++ // d = d + 1
	fmt.Println(d)
}
