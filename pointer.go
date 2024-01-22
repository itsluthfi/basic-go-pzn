package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// secara default, golang itu pass by value
	/**
	Address1   dicopy -->		Address2
		|							|
		|			-->				|
		|			-->				|
	address1					address2
	*/
	address1 := Address{"Bantul", "DI Yogyakarta", "Indonesia"}
	address2 := address1 // value dari address1 dicopy ke address2, value address1 != address2

	address2.City = "Sleman" // perubahan address2 tidak memengaruhi address1
	fmt.Println(address1)    // Sleman ga berubah karena ini value ori; out: {Bantul DI Yogyakarta Indonesia}
	fmt.Println(address2)    // berubah karena nilainya hasil copy dari address1; out: {Sleman DI Yogyakarta Indonesia}

	// pass by reference pake pointer
	/**
	Address1   dicopy -->	Address2		kalo pake pointer jadi gini					Address1
		|						|			-->											/	\
		|						|			--> 									/			\
		|						|												/					\
	address1				address2										address1			address2
	*/
	address3 := &address1 // operator & berarti dia nunjuk ke arah address1, jadi value address3 == address1
	/**
	kalo ditulis lengkap jadi gini
	var address1 Address = Address{"Bantul", "DI Yogyakarta", "Indonesia"}
	var address3 *Address = &address1
	tanda * berarti dia nunjuk ke Address yang udah ada, bukan bikin Address baru
	*/
	address3.City = "Kulonprogo"
	fmt.Println(address1) // ini ikut berubah; out: {Kulonprogo DI Yogyakarta Indonesia}
	fmt.Println(address3) // ini berubah jadi Kulonprogo; out: &{Kulonprogo DI Yogyakarta Indonesia}
}
