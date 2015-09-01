#ifndef GOPYTHON_GOPYTHON_H
#define GOPYTHON_GOPYTHON_H 1

#include "Python.h"

#include "frameobject.h"
#include "marshal.h"

/* stdlib */
#include <stdlib.h>
#include <string.h>

/* go-python */
#define _gopy_max_varargs 8 /* maximum number of varargs accepted by go-python */

/* --- object --- */

int
_gopy_PyObject_DelAttr(PyObject *o, PyObject *attr_name);

int
_gopy_PyObject_DelAttrString(PyObject *o, const char *attr_name);

PyObject*
_gopy_PyObject_CallFunction(PyObject *o, int len, char* types, void *args);

PyObject*
_gopy_PyObject_CallMethod(PyObject *o, char *method, int len, char* types, void *args);

/* --- dict --- */

int
_gopy_PyDict_Check(PyObject *o);

int
_gopy_PyDict_CheckExact(PyObject *o);

/* --- exceptions --- */

int
_gopy_PyErr_WarnPy3k(char *message, int stacklevel);

// FIXME we should instead have a exceptions_windows.go file...
# ifndef WIN32
PyObject* PyErr_SetFromWindowsErr(int ierr);
PyObject* PyErr_SetExcFromWindowsErr(PyObject *type, int ierr);
PyObject* PyErr_SetFromWindowsErrWithFilename(int ierr, const char *filename);
PyObject* PyErr_SetExcFromWindowsErrWithFilename(PyObject *type, int ierr, char *filename);
# endif

/* --- heap --- */

PyObject*
_gopy_InitModule(const char* nm, PyMethodDef* methods);

PyObject*
_gopy_InitModule3(const char* nm, PyMethodDef* methods, const char *doc);

PyMethodDef*
_gopy_malloc_PyMethodDefArray(size_t n);

void
_gopy_set_PyMethodDef(PyMethodDef *array, int i, PyMethodDef *o);

/* --- none --- */

PyObject*
_gopy_pynone(void);

/* --- numeric --- */

int
_gopy_PyInt_Check(PyObject *o);

int
_gopy_PyInt_CheckExact(PyObject *o);

long
_gopy_PyInt_AS_LONG(PyObject *io);

int
_gopy_PyLong_Check(PyObject *o);

int
_gopy_PyLong_CheckExact(PyObject *o);

int
_gopy_PyBool_Check(PyObject *o);

PyObject*
_gopy_pyfalse(void);

PyObject*
_gopy_pytrue(void);

int
_gopy_PyFloat_Check(PyObject *o);

int
_gopy_PyFloat_CheckExact(PyObject *o);

double
_gopy_PyFloat_AS_DOUBLE(PyObject *pyfloat);

int
_gopy_PyComplex_Check(PyObject *o);

int
_gopy_PyComplex_CheckExact(PyObject *o);

/* --- otherobjects --- */

int
_gopy_PyModule_Check(PyObject *p);

int
_gopy_PyModule_CheckExact(PyObject *p);

int
_gopy_PyClass_Check(PyObject *o);

int
_gopy_PyInstance_Check(PyObject *obj);

int
_gopy_PyFunction_Check(PyObject *o);

int
_gopy_PyMethod_Check(PyObject *o);

PyObject*
_gopy_PyMethod_GET_CLASS(PyObject *meth);

PyObject*
_gopy_PyMethod_GET_FUNCTION(PyObject *meth);

PyObject*
_gopy_PyMethod_GET_SELF(PyObject *meth);

int
_gopy_PySlice_Check(PyObject *ob);

int
_gopy_PyCapsule_CheckExact(PyObject *p);

int
_gopy_PyGen_Check(PyObject *ob);

int
_gopy_PyGen_CheckExact(PyObject *ob);

int
_gopy_PySeqIter_Check(PyObject *op);

int
_gopy_PyCallIter_Check(PyObject *op);

/* --- sequence --- */

int
_gopy_PyByteArray_Check(PyObject *o);

int
_gopy_PyByteArray_CheckExact(PyObject *o);

char*
_gopy_PyByteArray_AS_STRING(PyObject *bytearray);

Py_ssize_t
_gopy_PyByteArray_GET_SIZE(PyObject *bytearray);

int
_gopy_PyTuple_Check(PyObject *o);

int
_gopy_PyTuple_CheckExact(PyObject *o);

Py_ssize_t
_gopy_PyTuple_GET_SIZE(PyObject *p);

void
_gopy_PyTuple_SET_ITEM(PyObject *p, Py_ssize_t pos, PyObject *o);

PyObject*
_gopy_PyTuple_GET_ITEM(PyObject *p, Py_ssize_t pos);

int
_gopy_PyList_Check(PyObject *o);

int
_gopy_PyList_CheckExact(PyObject *o);

Py_ssize_t
_gopy_PyList_GET_SIZE(PyObject *o);

PyObject*
_gopy_PyList_GET_ITEM(PyObject *list, Py_ssize_t i);

void
_gopy_PyList_SET_ITEM(PyObject *list, Py_ssize_t i, PyObject *o);

int
_gopy_PyString_Check(PyObject *o);

Py_ssize_t
_gopy_PyString_GET_SIZE(PyObject *o);

char*
_gopy_PyString_AS_STRING(PyObject *o);

int
_gopy_PyObject_CheckBuffer(PyObject *obj);

int
_gopy_PyMemoryView_Check(PyObject *obj);

Py_buffer*
_gopy_PyMemoryView_GET_BUFFER(PyObject *obj);


/* --- type --- */

int
_gopy_PyType_Check(PyObject *o);

int
_gopy_PyType_CheckExact(PyObject *o);


/* --- utilities --- */

int
_gopy_PyOS_CheckStack(void);

PyObject*
_gopy_PyImport_ImportModuleEx(char *name, PyObject *globals, PyObject *locals, PyObject *fromlist);

//void
//_gopy_PySys_WriteStdout(const char *data);

/* --- veryhigh --- */

int
_gopy_PyRun_SimpleString(const char *command);

#endif /* !GOPYTHON_GOPYTHON_H */
