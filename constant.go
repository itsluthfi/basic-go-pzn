package main

func main() {
	const firstName string = "Luthfi"
	const lastName = "Izzuddin"
	// constant kalo ga dipake tidak masalah, tp gabisa diubah

	// error
	firstName = "Hanif"
	lastName = "Luthfi"

	// multiple constant
	const (
		myName   string = "Luthfi"
		yourName        = "Gatau"
	)
}
