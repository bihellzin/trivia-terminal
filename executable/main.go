package main

import (
	"connection/apiConnection"
	"fmt"
	"term/front"
)

func main() {
	test := apiConnection.GenerateAPIURL()
	fmt.Println(test)
	front.Run()
}
