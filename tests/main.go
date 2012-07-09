package main

import "fmt"
import "github.com/sbinet/go-python/pkg/python"

func main() {
	gostr := "foo"
	pystr := python.PyString_FromString(gostr)
	str := python.PyString_AsString(pystr)
	fmt.Println("hello [", str, "]")
}
