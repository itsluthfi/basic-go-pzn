package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // jadi -32767 karena ada numbe overflow, balik lagi ke nilai plg kecil

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name string = "Luthfi Izzuddin Hanif"
	var l uint8 = name[0]
	var lString = string(l)

	fmt.Println(name)
	fmt.Println(l)
	fmt.Println(lString)
}
