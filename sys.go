package python

// SetStdinObject sets a sys.stdin to a specified python object.
func (py *Interpreter) SetStdinObject(obj *Object) error {
	ret := py.r.SysSetObject("stdin", obj.ptr)
	return errCode(ret)
}

// SetStdoutObject sets a sys.stdout to a specified python object.
func (py *Interpreter) SetStdoutObject(obj *Object) error {
	ret := py.r.SysSetObject("stdout", obj.ptr)
	return errCode(ret)
}

// SetStderrObject sets a sys.stderr to a specified python object.
func (py *Interpreter) SetStderrObject(obj *Object) error {
	ret := py.r.SysSetObject("stderr", obj.ptr)
	return errCode(ret)
}
