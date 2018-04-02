package main

import (
	"fmt"
	"log"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Printf("importing kwargs...\n")
	m := python.PyImport_ImportModule("kwargs")
	if m == nil {
		log.Fatalf("could not import 'kwargs'\n")
	}

	foo := m.GetAttrString("foo")
	if foo == nil {
		log.Fatalf("could not getattr(kwargs, 'foo')\n")
	}

	out := foo.CallFunction()
	if out == nil {
		log.Fatalf("error calling foo()\n")
	}
	str := python.PyString_AsString(out)
	fmt.Printf("%s\n", str)
	want := "args=() kwds={}"
	if str != want {
		log.Fatalf("error. got=%q want=%q\n", str, want)
	}

	// keyword arguments
	kw := python.PyDict_New()
	err := python.PyDict_SetItem(
		kw,
		python.PyString_FromString("a"),
		python.PyInt_FromLong(3),
	)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	args := python.PyTuple_New(0)
	out = foo.Call(args, kw)
	if out == nil {
		log.Fatalf("error calling foo(*args, **kwargs)\n")
	}

	str = python.PyString_AsString(out)
	fmt.Printf("%s\n", str)
	want = "args=() kwds={'a': 3}"
	if str != want {
		log.Fatalf("error. got=%q. want=%q\n", str, want)
	}
}
