package main

import (
	"fmt"
	"log"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	module := python.PyImport_ImportModule("get_none")
	if module == nil {
		log.Fatal("could not import 'get_none' module")
	}
	get_none := module.GetAttrString("get_none")
	if get_none == nil {
		log.Fatal("could not import 'get_none' function")
	}
	none := get_none.CallFunction()
	fmt.Printf("type=%s, str=%s, eq_none=%t\n",
		python.PyString_AsString(none.Type().Str()),
		python.PyString_AsString(none.Str()),
		none == python.Py_None,
	)
}
