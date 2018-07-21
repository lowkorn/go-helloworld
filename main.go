package main

import (
	"fmt"

	"github.com/lowkorn/gotools"
	"github.com/lowkorn/gotools/stringhelper"
)

func main() {
	fmt.Println("h	ello")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("Dog"))
	fmt.Println(stringhelper.Concat("Hello", "Go"))
}
