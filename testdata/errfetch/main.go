package main

import (
	"fmt"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	exc, val, tb := python.PyErr_Fetch()
	fmt.Printf("exc=%v\nval=%v\ntb=%v\n", exc, val, tb)
}
