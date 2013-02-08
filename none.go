package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//PyObject *_gopy_pynone(void) { return Py_None; }
import "C"

// The Python None object, denoting lack of value. This object has no methods.
// It needs to be treated just like any other object with respect to reference
// counts.
var Py_None = togo(C._gopy_pynone())

// EOF
