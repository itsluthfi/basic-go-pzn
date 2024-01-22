package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

// struct method
func (person Person) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", person.Name)
}

func main() {
	luthfi := Person{}
	luthfi.Name = "Luthfi"
	luthfi.Address = "Indonesia"
	luthfi.Age = 22
	fmt.Println(luthfi)

	// struct literals
	izzuddin := Person{
		Name:    "Izzuddin",
		Address: "Yogyakarta",
		Age:     21,
	}
	fmt.Println(izzuddin)

	hanif := Person{"Hanif", "Tegalrejo", 23}
	fmt.Println(hanif)

	luthfi.sayHello("Eko")
	izzuddin.sayHello("Eko")
	hanif.sayHello("Eko")
}
