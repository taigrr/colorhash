package main

import (
	"fmt"

	go_colorhash "github.com/taigrr/colorhash"
)

func main() {
	x := go_colorhash.HashString("asdasd")
	fmt.Println(x)
}
