package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// bikin pointer pake new, tapi datanya kosong gaada isinya

	// tanpa new
	var alamat1 *Address = &Address{}
	var alamat2 *Address = alamat1

	alamat2.City = "Yogyakarta"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

	// pake new
	var alamat3 *Address = new(Address)
	var alamat4 *Address = alamat3

	alamat4.Country = "Indonesia"

	fmt.Println(alamat3)
	fmt.Println(alamat4)
}
