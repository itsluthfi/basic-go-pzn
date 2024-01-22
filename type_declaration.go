package main

import "fmt"

func main() {
	type NoKTP string

	var ktpLuthfi NoKTP = "1111111111"

	var contoh string = "2222222222"
	var ktpContoh NoKTP = NoKTP(contoh)

	fmt.Println(ktpLuthfi)
	fmt.Println(ktpContoh)
}
