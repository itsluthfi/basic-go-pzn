package main

import "fmt"

func main() {
	name := "Hanif"

	if name == "Luthfi" {
		fmt.Println("Halo Luthfi")
	} else if name == "Izzuddin" {
		fmt.Println("Halo Izzuddin")
	} else {
		fmt.Println("Halo, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
