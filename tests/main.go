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

	pickle := python.PyImport_ImportModule("cPickle")
	if pickle == nil {
		panic("could not import 'cPickle'")
	}
	dumps := pickle.GetAttrString("dumps")
	if dumps == nil {
		panic("could not retrieve 'cPickle.dumps'")
	}
	out := dumps.CallFunctionObjArgs("O", pystr)
	if out == nil {
		panic("could not dump pystr")
	}
	fmt.Printf("cPickle.dumps(%s) = %q\n", gostr,
		python.PyString_AsString(out),
	)
}
