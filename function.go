package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// function params
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// return value
func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

// return multiple values
func getFullName() (string, string) {
	return "Luthfi", "Izzuddin"
}

// named return values
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Luthfi"
	// middleName = "Izzuddin"
	lastName = "Hanif"

	return firstName, middleName, lastName
}

func main() {
	sayHello()

	sayHelloTo("Luthfi", "Izzuddin")

	getHello("Luthfi Izzuddin") // kosong karena nilainya tidak diassign ke variabel
	result := getHello("Luthfi Izzuddin")
	fmt.Println(result)
	// atau bisa gini kalo ga pengen assign di variabel
	fmt.Println(getHello("Luthfi"))
	fmt.Println(getHello("Izzuddin"))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	_, newLastName := getFullName()
	fmt.Println(newLastName)

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
