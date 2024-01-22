package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := pembagian(110, 0)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("error", err.Error())
	}
}
