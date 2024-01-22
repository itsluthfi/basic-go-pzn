package main

import "fmt"

func main() {
	// var car map[string]string
	// bike := map[string]string{}

	person := map[string]string{
		"nama":   "Luthfi",
		"alamat": "Tegalrejo",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])

	book := make(map[string]string)
	book["title"] = "Belajar Go Lang"
	book["author"] = "Luthfi"
	book["tahun"] = "2023"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}
