package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perlulangan ke", counter)
		counter++
	}

	for newCounter := 1; newCounter <= 10; newCounter++ {
		fmt.Println("Perulangan baru ke", newCounter)
	}

	// for range
	names := []string{"Luthfi", "Izzuddin", "Hanif"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "name =", name)
	}

	for _, name := range names {
		fmt.Println("Name =", name)
	}
}
