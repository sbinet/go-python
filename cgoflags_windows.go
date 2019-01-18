// +build windows

package python

// #cgo 386   CFLAGS: -I C:/Python27/include
// #cgo amd64 CFLAGS: -I C:/Python27-x64/include
// #cgo amd64 CFLAGS: -D MS_WIN64
// #cgo 386   LDFLAGS: -L C:/Python27/libs     -lpython27
// #cgo amd64 LDFLAGS: -L C:/Python27-x64/libs -lpython27
//
// #include "go-python.h"
import "C"
