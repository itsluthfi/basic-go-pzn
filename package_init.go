package main

import (
	"basic-go-pzn/database"
	_ "basic-go-pzn/internal" // kalo mau jalanin init functionnya aja tanpa func yg lain, ditambahin blank identifier (_) sebelum import
	"fmt"
)

func main() {
	// output pertama adalah hasil dari fungsi init di package internal

	result := database.GetDatabase()
	fmt.Println(result)
}
