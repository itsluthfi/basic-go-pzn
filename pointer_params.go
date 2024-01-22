package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
	fmt.Println(address)
}

func PointerChangeCountryToIndonesia(address *Address) { // ditambah * di bagian tipe data paramsnya
	address.Country = "Indonesia"
	fmt.Println(address)
}

func main() {
	// secara default, params di fungsi itu juga pass by value, jadi dicopy dulu
	address := Address{}

	ChangeCountryToIndonesia(address) // ada isi Indonesianya di akhir, karena address dicopy dulu ke dalem fungsinya
	fmt.Println(address)              // (variabel asli) outputnya aslinya tetep kosong karena tadi dicopy dulu

	PointerChangeCountryToIndonesia(&address) // paramsnya ditambahin &, atau kalo gamau ditambahin & harus bikin variabel tsb pointer dari awal
	fmt.Println(address)

	// pointer dari awal
	address1 := &Address{} // ini langsung pointer
	PointerChangeCountryToIndonesia(address1)
	fmt.Println(address)
}
