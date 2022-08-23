package main

import (
	"fmt"

	"githum.com/tvn9/gomodules/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random string:", s)
}
