package main

import (
	"basic-go-pzn/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Luthfi")
	fmt.Println(result)

	fmt.Println(helper.Application)
}
