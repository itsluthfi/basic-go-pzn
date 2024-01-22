package main

import "fmt"

type HasName interface { // interface hanya khusus method, method = function yang ada dlm struct/interface
	GetName() string
}

func SayHello(hasName HasName) { // terafiliasi sama interface
	fmt.Println("Hello", hasName.GetName())
}

type People struct {
	Name string
}

func (people People) GetName() string { // terafiliasi sama People struct dan mengimplementasikan interface HasName, berarti bisa pake SayHello, karena ada fungsi GetName dan returnnya string
	return people.Name
}

// interface kosong/any
func random() interface{} {
	return 1
}

func main() {
	hanif := People{"Hanif"}

	SayHello(hanif)

	var random interface{} = random()
	// random := random()

	fmt.Println(random)
}
