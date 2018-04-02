package python

func (py *Interpreter) setStdStreamObject(name string, obj *Object) error {
	ret := py.r.SysSetObject(name, obj.ptr)
	if ret != 0 {
		return errCode(ret)
	}
	ret = py.r.SysSetObject("__"+name+"__", obj.ptr)
	return errCode(ret)
}

// SetStdinObject sets a sys.stdin to a specified python object.
func (py *Interpreter) SetStdinObject(obj *Object) error {
	return py.setStdStreamObject("stdin", obj)
}

// SetStdoutObject sets a sys.stdout to a specified python object.
func (py *Interpreter) SetStdoutObject(obj *Object) error {
	return py.setStdStreamObject("stdout", obj)
}

// SetStderrObject sets a sys.stderr to a specified python object.
func (py *Interpreter) SetStderrObject(obj *Object) error {
	return py.setStdStreamObject("stderr", obj)
}
