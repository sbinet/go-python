// a go wrapper around py-main
package main

import (
	"github.com/sbinet/go-python"
	"os"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	rc := python.Py_Main(os.Args)
	os.Exit(rc)
}
