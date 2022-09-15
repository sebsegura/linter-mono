package main

import (
	"fmt"
	"sebsegura/linter-mono/pkg/adder"
)

const ID = "example"

func main() {
	fmt.Println(ID)
	fmt.Println(adder.Add(2, 2))
}
