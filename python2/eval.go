package python2

/*
#include "go-python.h"

// Forward declaration.
int go_trace_cgo(PyObject *obj, PyFrameObject *frame, int what, PyObject *arg);
*/
import "C"

import (
	"github.com/sbinet/go-python/runtime"
	"unsafe"
)

type Frame struct {
	ptr *C.PyFrameObject
}

func (f *Frame) toPy() *C.PyFrameObject {
	if f == nil {
		return nil
	}
	return f.ptr
}
func (f *Frame) GetLineNumber() int {
	return int(C.PyFrame_GetLineNumber(f.ptr))
}
func (f *Frame) GetFilename() runtime.Object {
	if f.ptr.f_code == nil {
		return nil
	}
	c := f.ptr.f_code
	return toGo(c.co_filename)
}
func (f *Frame) Builtins() runtime.DictObject {
	if f.ptr == nil {
		return nil
	}
	return toGo(f.ptr.f_builtins)
}
func (f *Frame) Globals() runtime.DictObject {
	if f.ptr == nil {
		return nil
	}
	return toGo(f.ptr.f_globals)
}
func (f *Frame) Locals() runtime.Object {
	if f.ptr == nil {
		return nil
	}
	return toGo(f.ptr.f_locals)
}

func (py2Runtime) EvalThreadsInitialized() bool {
	return C.PyEval_ThreadsInitialized() != 0
}

func (py2Runtime) EvalInitThreads() {
	C.PyEval_InitThreads()
}

func (py2Runtime) EvalGetFrame() runtime.Frame {
	f := C.PyEval_GetFrame()
	if f == nil {
		return nil
	}
	return &Frame{ptr: f}
}

var (
	traces     = make(map[*C.PyObject]runtime.TraceFunc)
	traceTypes = map[int]runtime.TraceType{
		int(C.PyTrace_CALL):        runtime.TraceCall,
		int(C.PyTrace_EXCEPTION):   runtime.TraceException,
		int(C.PyTrace_LINE):        runtime.TraceLine,
		int(C.PyTrace_RETURN):      runtime.TraceReturn,
		int(C.PyTrace_C_CALL):      runtime.TraceCCall,
		int(C.PyTrace_C_EXCEPTION): runtime.TraceCException,
		int(C.PyTrace_C_RETURN):    runtime.TraceCReturn,
	}
)

//export go_trace
func go_trace(obj *C.PyObject, frame *C.PyFrameObject, what int, arg *C.PyObject) int {
	fnc := traces[obj]
	tt := traceTypes[what]
	return fnc(toGo(obj), &Frame{ptr: frame}, tt, toGo(arg))
}

func (py2Runtime) EvalSetTrace(fnc runtime.TraceFunc, obj runtime.Object) {
	o := fromPtr(obj).toPy()
	traces[o] = fnc
	C.PyEval_SetTrace((C.Py_tracefunc)(unsafe.Pointer(C.go_trace_cgo)), o)
}
