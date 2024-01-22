package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}

	for idx := 0; idx < 10; idx++ {
		if idx%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke-", idx)
	}
}
