package python

// #include "go-python.h"
import "C"

// PyObject* Py_BuildValue(const char *format, ...)
// Return value: New reference.
func Py_BuildValue(format string, args ...interface{}) *PyObject {
	return nil
}

// PyMethodDef
// ml_name	char *	name of the method
// ml_meth	PyCFunction	pointer to the C implementation
// ml_flags	int	flag bits indicating how the call should be constructed
// ml_doc	char *	points to the contents of the docstring
type PyMethodDef struct {
	Name  string // name of the method
	Meth  PyCFunction
	Flags MethodDefFlags
	Doc   string
}

type PyCFunction C.PyCFunction

type MethodDefFlags int

const (
	MethVarArgs  MethodDefFlags = C.METH_VARARGS
	MethKeyWords                = C.METH_KEYWORDS
	MethNoArgs                  = C.METH_NOARGS
	MethO                       = C.METH_O
	MethOldArgs                 = C.METH_OLDARGS
	MethClass                   = C.METH_CLASS
	MethStatic                  = C.METH_STATIC
	MethCoexist                 = C.METH_COEXIST
)
