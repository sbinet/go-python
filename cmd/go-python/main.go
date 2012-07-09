// a go wrapper around py-main
package main

import "os"
import "github.com/sbinet/go-python/pkg/python"

func main() {
	rc := python.Py_Main(os.Args)
	os.Exit(rc)
}
