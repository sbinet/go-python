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
	gostr := "foo"
	pystr := python.PyString_FromString(gostr)
	str := python.PyString_AsString(pystr)
	fmt.Println("hello [", str, "]")
}
