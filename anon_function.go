package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Nama Anda diblacklist,", name)
	} else {
		fmt.Println("Welcome,", name)
	}

}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", blacklist)

	registerUser("Luthfi", func(name string) bool {
		return name == "Anjing"
	})
}
