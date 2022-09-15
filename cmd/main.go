package main

import (
	"fmt"
	"sebsegura/linter-mono/pkg/adder"
)

const Id = "example"

func main() {
	fmt.Println(Id)
	fmt.Println(adder.Add(2, 2))
}
