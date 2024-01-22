package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Luthfi"
	names[1] = "Izzuddin"
	names[2] = "Hanif"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		85,
		80,
		65,
		75,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
