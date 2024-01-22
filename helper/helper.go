package helper

/*
access modifier

secara default, kalo ada variable, function, method pake huruf kecil di awal, berarti dia cuman bisa diakses di package itu aja

kalo pake huruf kapital, berarti dia bisa diakses di luar package
**/

var version = "1.0"        // gabisa diakses di luar package
var Application = "golang" // bisa diakses di luar package helper

func SayHello(name string) string { // bisa diakses di luar package helper
	return "Hello " + name
}

func sayGoodBye(name string) string {
	return "Good bye " + name
}

// walaupun gabisa diakses di luar package, tapi itu tetep bisa dipake di package tsb
