package python2

/*
#include "go-python.h"

// The gateway function
int go_trace_cgo(PyObject *obj, PyFrameObject *frame, int what, PyObject *arg) {
	int go_trace(PyObject *obj, PyFrameObject *frame, int what, PyObject *arg);
	return go_trace(obj, frame, what, arg);
}
*/
import "C"
