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

	origList := python.PyList_New(0)
	python.PyList_Append(origList, python.PyString_FromString("i want this gone"))
	fmt.Println(python.PyString_AsString(origList.Str()))

	python.PyList_SetSlice(origList, 0, 1, nil)
	fmt.Println(python.PyString_AsString(origList.Str()))
}
