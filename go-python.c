#include "go-python.h"

/* --- object --- */

int
_gopy_PyObject_DelAttr(PyObject *o, PyObject *attr_name) {
	return PyObject_DelAttr(o, attr_name);
}

int
_gopy_PyObject_DelAttrString(PyObject *o, const char *attr_name) {
	return PyObject_DelAttrString(o,attr_name);
}

PyObject*
_gopy_PyObject_CallFunction(PyObject *o, int len, char* pyfmt, void *cargs) {
	void ** args = (void**)cargs;

	if (len > _gopy_max_varargs) {
			PyErr_Format(
					PyExc_RuntimeError, 
					"python: maximum number of varargs (%d) exceeded (%d)",
					_gopy_max_varargs,
					len
			);
			return NULL;
	}

	switch (len) {
		case 0:
			return PyObject_CallFunction(o, pyfmt);

		case 1:
			return PyObject_CallFunction(o, pyfmt, args[0]);

		case 2:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1]);

		case 3:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2]);

		case 4:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3]);

		case 5:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4]);

		case 6:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5]);

		case 7:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5], args[6]);

		case 8:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7]);

		default:
			PyErr_Format(PyExc_RuntimeError, "python: invalid number of arguments (%d)", len);
			return NULL;

	}

	return NULL;
}

PyObject*
_gopy_PyObject_CallMethod(PyObject *o, char *method, int len, char* pyfmt, void *cargs) {
	void ** args = (void**)cargs;

	if (len > _gopy_max_varargs) {
			PyErr_Format(
					PyExc_RuntimeError, 
					"python: maximum number of varargs (%d) exceeded (%d)",
					_gopy_max_varargs,
					len
			);
			return NULL;
	}

	switch (len) {
		case 0:
			return PyObject_CallMethod(o, method, pyfmt);

		case 1:
			return PyObject_CallMethod(o, method, pyfmt, args[0]);

		case 2:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1]);

		case 3:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1], args[2]);

		case 4:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1], args[2], args[3]);

		case 5:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1], args[2], args[3], args[4]);

		case 6:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5]);

		case 7:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5], args[6]);

		case 8:
			return PyObject_CallMethod(o, method, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7]);

		default:
			PyErr_Format(PyExc_RuntimeError, "python: invalid number of arguments (%d)", len);
			return NULL;

	}

	return NULL;
}


/* --- dict --- */

int _gopy_PyDict_Check(PyObject *o) { 
	return PyDict_Check(o); 
}

int _gopy_PyDict_CheckExact(PyObject *o) { 
	return PyDict_CheckExact(o); 
}

/* --- exceptions --- */

int
_gopy_PyErr_WarnPy3k(char *message, int stacklevel) {
	return PyErr_WarnPy3k(message, stacklevel);
}

// FIXME we should instead have a exceptions_windows.go file...
# ifndef WIN32
PyObject* PyErr_SetFromWindowsErr(int ierr) { return NULL; }
PyObject* PyErr_SetExcFromWindowsErr(PyObject *type, int ierr) { return NULL; }
PyObject* PyErr_SetFromWindowsErrWithFilename(int ierr, const char *filename) { return NULL; }
PyObject* PyErr_SetExcFromWindowsErrWithFilename(PyObject *type, int ierr, char *filename) { return NULL; }
# endif

/* --- heap --- */

PyObject*
_gopy_InitModule(const char* nm, PyMethodDef* methods) {
	return Py_InitModule(nm, methods);
}

PyObject*
_gopy_InitModule3(const char* nm, PyMethodDef* methods, const char *doc) {
	return Py_InitModule3(nm, methods, doc);
}

PyMethodDef*
_gopy_malloc_PyMethodDefArray(size_t n) {
    return (PyMethodDef*)malloc(n * sizeof(PyMethodDef));
}

void
_gopy_set_PyMethodDef(PyMethodDef *array, int i, PyMethodDef *o) {
    array[i] = *o;
}

/* --- none --- */

PyObject*
_gopy_pynone(void) { 
	return Py_None; 
}

/* --- numeric --- */

int
_gopy_PyInt_Check(PyObject *o) { 
	return PyInt_Check(o);
}

int
_gopy_PyInt_CheckExact(PyObject *o) { 
	return PyInt_CheckExact(o);
}

long
_gopy_PyInt_AS_LONG(PyObject *o) { 
	return PyInt_AS_LONG(o);
}

int
_gopy_PyLong_Check(PyObject *o) { 
	return PyLong_Check(o);
}

int _gopy_PyLong_CheckExact(PyObject *o) { 
	return PyLong_CheckExact(o);
}

int
_gopy_PyBool_Check(PyObject *o) {
	return PyBool_Check(o);
}

PyObject*
_gopy_pyfalse(void) {
	return Py_False;
}

PyObject*
_gopy_pytrue(void) {
	return Py_True;
}

int
_gopy_PyFloat_Check(PyObject *o) {
	return PyFloat_Check(o);
}

int
_gopy_PyFloat_CheckExact(PyObject *o) {
	return PyFloat_CheckExact(o);
}

double
_gopy_PyFloat_AS_DOUBLE(PyObject *pyfloat) {
	return PyFloat_AS_DOUBLE(pyfloat);
}

int
_gopy_PyComplex_Check(PyObject *o) {
	return PyComplex_Check(o);
}

int
_gopy_PyComplex_CheckExact(PyObject *o) {
	return PyComplex_CheckExact(o);
}

/* --- otherobjects --- */

int
_gopy_PyModule_Check(PyObject *p) { 
	return PyModule_Check(p);
}

int
_gopy_PyModule_CheckExact(PyObject *p) {
	return PyModule_CheckExact(p);
}

int
_gopy_PyClass_Check(PyObject *o) {
	return PyClass_Check(o);
}

int
_gopy_PyInstance_Check(PyObject *obj) {
	return PyInstance_Check(obj);
}

int
_gopy_PyFunction_Check(PyObject *o) {
	return PyFunction_Check(o);
}

int
_gopy_PyMethod_Check(PyObject *o) {
	return PyMethod_Check(o);
}

PyObject*
_gopy_PyMethod_GET_CLASS(PyObject *meth) {
	return PyMethod_GET_CLASS(meth);
}

PyObject*
_gopy_PyMethod_GET_FUNCTION(PyObject *meth) {
	return PyMethod_GET_FUNCTION(meth);
}

PyObject*
_gopy_PyMethod_GET_SELF(PyObject *meth) {
	return PyMethod_GET_SELF(meth);
}

int
_gopy_PySlice_Check(PyObject *ob) {
	return PySlice_Check(ob);
}

int
_gopy_PyCapsule_CheckExact(PyObject *p) {
	return PyCapsule_CheckExact(p);
}

int
_gopy_PyGen_Check(PyObject *ob) {
	return _gopy_PyGen_Check(ob);
}

int
_gopy_PyGen_CheckExact(PyObject *ob) {
	return _gopy_PyGen_CheckExact(ob);
}

int
_gopy_PySeqIter_Check(PyObject *op) {
	return PySeqIter_Check(op);
}

int
_gopy_PyCallIter_Check(PyObject *op) {
	return PyCallIter_Check(op);
}

/* --- sequence --- */

int
_gopy_PyByteArray_Check(PyObject *o) {
	return PyByteArray_Check(o);
}

int
_gopy_PyByteArray_CheckExact(PyObject *o) {
	return PyByteArray_CheckExact(o);
}

char*
_gopy_PyByteArray_AS_STRING(PyObject *bytearray) {
	return PyByteArray_AS_STRING(bytearray);
}

Py_ssize_t
_gopy_PyByteArray_GET_SIZE(PyObject *bytearray) {
	return PyByteArray_GET_SIZE(bytearray);
}

int
_gopy_PyTuple_Check(PyObject *o) {
	return PyTuple_Check(o);
}

int
_gopy_PyTuple_CheckExact(PyObject *o) {
	return PyTuple_CheckExact(o);
}

Py_ssize_t
_gopy_PyTuple_GET_SIZE(PyObject *p) {
	return PyTuple_GET_SIZE(p);
}

void
_gopy_PyTuple_SET_ITEM(PyObject *p, Py_ssize_t pos, PyObject *o) {
	PyTuple_SET_ITEM(p, pos, o);
}

PyObject*
_gopy_PyTuple_GET_ITEM(PyObject *p, Py_ssize_t pos) {
	return PyTuple_GET_ITEM(p, pos);
}

int
_gopy_PyList_Check(PyObject *o) {
	return PyList_Check(o);
}

int
_gopy_PyList_CheckExact(PyObject *o) {
	return PyList_CheckExact(o);
}

Py_ssize_t
_gopy_PyList_GET_SIZE(PyObject *o) {
	return PyList_GET_SIZE(o);
}

PyObject*
_gopy_PyList_GET_ITEM(PyObject *list, Py_ssize_t i) {
	return PyList_GET_ITEM(list, i);
}

void
_gopy_PyList_SET_ITEM(PyObject *list, Py_ssize_t i, PyObject *o) {
	PyList_SET_ITEM(list, i, o);
}

int
_gopy_PyString_Check(PyObject *o) {
	return PyString_Check(o);
}

Py_ssize_t
_gopy_PyString_GET_SIZE(PyObject *o) { 
	return PyString_GET_SIZE(o);
}

char*
_gopy_PyString_AS_STRING(PyObject *o) {
	return PyString_AS_STRING(o);
}

int
_gopy_PyObject_CheckBuffer(PyObject *obj) {
	return PyObject_CheckBuffer(obj);
}

int
_gopy_PyMemoryView_Check(PyObject *obj) {
	return PyMemoryView_Check(obj);
}

Py_buffer*
_gopy_PyMemoryView_GET_BUFFER(PyObject *obj) {
	return PyMemoryView_GET_BUFFER(obj);
}

/* --- type --- */

int
_gopy_PyType_Check(PyObject *o) {
	return PyType_Check(o);
}

int
_gopy_PyType_CheckExact(PyObject *o) {
	return PyType_CheckExact(o);
}


/* --- utilities --- */

#ifdef USE_STACKCHECK
 int
 _gopy_PyOS_CheckStack(void) {
	 return PyOS_CheckStack();
 }

#else
 int
 _gopy_PyOS_CheckStack(void) {
	 return 0;
 }
#endif

PyObject*
_gopy_PyImport_ImportModuleEx(char *name, PyObject *globals, PyObject *locals, PyObject *fromlist) {
	return PyImport_ImportModuleEx(name, globals, locals, fromlist);
}

//void _gopy_PySys_WriteStdout(const char *data) {
// PySys_WriteStdout("%s", data)
//}

/* --- veryhigh --- */

int
_gopy_PyRun_SimpleString(const char *command) {
	return PyRun_SimpleString(command);
}
