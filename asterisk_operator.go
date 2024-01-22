package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Address{"Bantul", "DI Yogyakarta", "Indonesia"}
	address2 := &address1

	address2.City = "Kulonprogo"
	fmt.Println(address1) // out: {Kulonprogo DI Yogyakarta Indonesia}
	fmt.Println(address2) // out: &{Kulonprogo DI Yogyakarta Indonesia}

	// memisahkan pointer address2 ke address lain, tapi address1 tetep valuenya
	/**
				Address1						Address1			Address2
				/	\				-->				|					|
			/			\			-->				|					|
		/					\						|					|
	address1			address2				address1			address2
	*/

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // out: {Kulonprogo DI Yogyakarta Indonesia}
	fmt.Println(address2) // out: &{Jakarta DKI Jakarta Indonesia}

	// gampangnya ganti value address1 & address2 lewat address2
	/**
				Address1						Address1			Address2
				/	\				-->								/	\
			/			\			-->							/			\
		/					\								/					\
	address1			address2						address1			address2
	*/

	*address2 = Address{"Semarang", "Jawa Tengah", "Indonesia"}
	fmt.Println(address1) // out: {Kulonprogo DI Yogyakarta Indonesia}
	fmt.Println(address2) // out: &{Jakarta DKI Jakarta Indonesia}
}
