package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++ // !hati-hati ini bakal nambahin nilai di variabel counter di luar fungsinya
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
