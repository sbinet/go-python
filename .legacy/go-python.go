package python

// #include "go-python.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// pyfmt returns the python format string for a given go value
func pyfmt(v interface{}) (unsafe.Pointer, string) {
	switch v := v.(type) {
	case bool:
		return unsafe.Pointer(&v), "b"

		// 	case byte:
		// 		return unsafe.Pointer(&v), "b"

	case int8:
		return unsafe.Pointer(&v), "b"

	case int16:
		return unsafe.Pointer(&v), "h"

	case int32:
		return unsafe.Pointer(&v), "i"

	case int64:
		return unsafe.Pointer(&v), "k"

	case int:
		switch unsafe.Sizeof(int(0)) {
		case 4:
			return unsafe.Pointer(&v), "i"
		case 8:
			return unsafe.Pointer(&v), "k"
		}

	case uint8:
		return unsafe.Pointer(&v), "B"

	case uint16:
		return unsafe.Pointer(&v), "H"

	case uint32:
		return unsafe.Pointer(&v), "I"

	case uint64:
		return unsafe.Pointer(&v), "K"

	case uint:
		switch unsafe.Sizeof(uint(0)) {
		case 4:
			return unsafe.Pointer(&v), "I"
		case 8:
			return unsafe.Pointer(&v), "K"
		}

	case float32:
		return unsafe.Pointer(&v), "f"

	case float64:
		return unsafe.Pointer(&v), "d"

	case complex64:
		return unsafe.Pointer(&v), "D"

	case complex128:
		return unsafe.Pointer(&v), "D"

	case string:
		cstr := C.CString(v)
		return unsafe.Pointer(cstr), "s"

	case *PyObject:
		return unsafe.Pointer(v.topy()), "O"

	}

	panic(fmt.Errorf("python: unknown type (%T)", v))
}
