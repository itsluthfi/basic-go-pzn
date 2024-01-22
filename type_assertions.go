package main

import (
	"fmt"
)

func Random() interface{} {
	return true
}

func main() {
	result := Random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// panic karena typenya bukan int
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// cara amannya pake switch
	switch value := result.(type) { // type assertions hanya untuk interface kosong
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
