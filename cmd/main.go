// a go wrapper around py-main
package main

import "os"
import "python"

func main() {
	rc := python.Main(os.Args)
	os.Exit(rc)
}