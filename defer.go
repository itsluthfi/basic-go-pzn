package main

import "fmt"

func logger() {
	fmt.Println("Dipanggil setelah selesai")
}

func runApp() {
	defer logger() // walaupun dipanggil di atas, bakal dijalanin setelah kode2 di bawahnya berjalan
	fmt.Println("Aplikasi jalan")
}

func main() {
	runApp()
}
