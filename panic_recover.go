package main

import "fmt"

func endApps() {
	fmt.Println("Program berhenti")
	message := recover() // kasih recover di dalam defer, kalo ga nanti gaakan muncul pesan panicnya, karena kalo panic nanti kode2 di bawahnya gaakan jalan
	fmt.Println(message)
}

func runApps(error bool) {
	defer endApps()
	if error {
		panic("Ups error")
	}

}

func main() {
	runApps(true)
	fmt.Println("Apakah berjalan?") // gaakan jalan kalo recovernya ga di dalem defer
}
