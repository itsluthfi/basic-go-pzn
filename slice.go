package main

import "fmt"

func main() {
	names := [...]string{"Luthfi", "Izzuddin", "Hanif", "Budi", "Joko", "Asep"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	var slice3 []string = names[2:5]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur Baru")
	fmt.Println(daysSlice2)

	daysSlice2[0] = "Sabtu Lama"
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice1 := make([]string, 2, 5)
	newSlice1[0] = "Luthfi"
	newSlice1[1] = "Izzuddin"
	// newSlice1[2] = "Hanif" // error karena lengthnya cuman 2, kalo mau nambah harus append

	fmt.Println(newSlice1)

	newSlice2 := append(newSlice1, "Hanif")
	fmt.Println(newSlice2)

	newSlice2[0] = "Lupek"
	fmt.Println(newSlice2)
	fmt.Println(newSlice1) // ikut keubah karena masih pake 1 array yg sama kayak newSlice1

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// ! hati hati kalo mau bikin array/slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	fmt.Println(len(iniSlice))
	fmt.Println(cap(iniSlice))
}
