package main

import "fmt"

func main() {
	newName := "Izzuddin"

	switch newName {
	case "Luthfi":
		fmt.Println("Halo Luthfi")
	case "Izzuddin":
		fmt.Println("Halo Izzuddin")
	default:
		fmt.Println("Halo, boleh kenalan?")
	}

	switch newLength := len(newName); newLength > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(newName)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
