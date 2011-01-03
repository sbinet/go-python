package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//PyObject *_gopy_pynone() { return Py_None; }
import "C"

var Py_None = togo(C._gopy_pynone())

// EOF
