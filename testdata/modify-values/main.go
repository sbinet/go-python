package main

import (
	"fmt"
	"log"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		log.Panic(err.Error())
	}
}

func main() {
	module := python.PyImport_ImportModule("values")
	if module == nil {
		log.Fatal("could not import 'values'")
	}

	name := module.GetAttrString("__name__")
	if name == nil {
		log.Fatal("could not getattr '__name__'")
	}
	defer name.DecRef()
	fmt.Printf("values.__name__: %q\n", python.PyString_AsString(name))

	sval := module.GetAttrString("sval")
	if sval == nil {
		log.Fatal("could not getattr 'sval'")
	}
	defer sval.DecRef()
	fmt.Printf("values.sval: %q\n", python.PyString_AsString(sval))

	pyival := module.GetAttrString("ival")
	if pyival == nil {
		log.Fatal("could not getattr 'ival'")
	}
	defer pyival.DecRef()

	ival := python.PyInt_AsLong(pyival)
	fmt.Printf("values.ival: %d\n", ival)

	myfunc := module.GetAttrString("myfunc")
	if myfunc == nil {
		log.Fatal("could not getattr 'myfunc'")
	}
	defer myfunc.DecRef()

	o1 := myfunc.CallFunction()
	if o1 == nil {
		log.Fatal("could not call 'values.myfunc()'")
	}
	fmt.Printf("%s\n", python.PyString_AsString(o1))
	o1.DecRef()

	// modify 'test.ival' and 'test.sval'
	{
		pyival := python.PyInt_FromLong(ival + 1000)
		module.SetAttrString("ival", pyival)
		pyival.DecRef()

		pysval := python.PyString_FromString(python.PyString_AsString(sval) + " is the answer")
		module.SetAttrString("sval", pysval)
		pysval.DecRef()
	}

	o2 := myfunc.CallFunction()
	if o2 == nil {
		log.Fatal("could not call 'values.myfunc()'")
	}
	fmt.Printf("%s\n", python.PyString_AsString(o2))
	o2.DecRef()
}
