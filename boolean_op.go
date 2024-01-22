package main

import "fmt"

func main() {
	var nilaiUjian = 90
	var presensi = 81

	var lulusNilaiUjian bool = nilaiUjian > 80
	var lulusPresensi bool = presensi > 80

	var lulus = lulusNilaiUjian && lulusPresensi

	fmt.Println(lulus)
}
