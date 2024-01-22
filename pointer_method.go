package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println(man)
}

func (man *Man) isMarried() {
	man.Name = "Mr. " + man.Name
	fmt.Println(man)
}

func main() {
	// method juga sama, secara default pass by value
	luthfi := Man{"Luthfi"}
	luthfi.Married() // harusnya di sini namanya jadi Mr. Luthfi

	fmt.Println(luthfi) // tapi di sini valuenya tetep Luthfi, agak aneh karena harusnya struct/object biasanya saling terkait

	// pake pointer biar data aslinya ikut keganti
	izzuddin := Man{"Izzuddin"}
	izzuddin.isMarried()

	fmt.Println(izzuddin)

	//* TIAP BIKIN METHOD SANGAT DIREKOMENDASIKAN SELALU PAKE POINTER --MAS EKO PZN
}
