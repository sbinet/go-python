// a go wrapper around py-main
package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/sbinet/go-python"
)

func loadRuntime(vers string) python.Runtime {
	pl, err := plugin.Open(fmt.Sprintf("python%s.so", vers))
	if err != nil {
		panic(err)
	}
	s, err := pl.Lookup("Runtime")
	if err != nil {
		panic(err)
	}
	r, ok := s.(*python.Runtime)
	if !ok {
		panic(fmt.Errorf("unexpected type: %T", s))
	}
	return *r
}

func main() {
	vers := "2"
	if os.Getenv("GO_PYTHON") == "3" {
		vers = "3"
	}
	r := loadRuntime(vers)
	py := python.NewInterpreter(r)
	err := py.Initialize()
	if err != nil {
		panic(err)
	}
	defer py.Close()

	err = py.Main(os.Args)
	if e, ok := err.(python.RunError); ok {
		os.Exit(e.Code)
	} else if err != nil {
		panic(err)
	}
}
