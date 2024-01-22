package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// *nilai nil cuman bisa ditaruh di tipe data interface, function, map, slice, pointer, dan channel

	NewMap("")

	data := NewMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
