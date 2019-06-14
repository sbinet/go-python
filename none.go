package python

// #include "go-python.h"
import "C"

// The Python None object, denoting lack of value. This object has no methods.
// It needs to be treated just like any other object with respect to reference
// counts.
var Py_None = &PyObject{ptr: C._gopy_pynone()}

// EOF
