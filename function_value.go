package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodBye := getGoodBye

	fmt.Println(goodBye("Luthfi"))
}
