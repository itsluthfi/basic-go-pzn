package database

var connection string

func init() { // dieksekusi pertama kali ketika diimport di package lain
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
