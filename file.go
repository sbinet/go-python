package python

import "os"

func (py *Interpreter) fromFile(f *os.File, mode string) *Object {
	p := py.r.FromFile(f, mode)
	return newObject(p)
}

// FromFile converts Go file into python file object.
func (py *Interpreter) FromFile(f *os.File, mode string) *Object {
	obj := py.fromFile(f, mode)
	obj.setFinalizer()
	return obj
}

// SetStdinFile sets a sys.stdin to a specified file descriptor.
func (py *Interpreter) SetStdinFile(f *os.File) error {
	pf := py.fromFile(f, "r")
	defer pf.decRef()
	return py.SetStdinObject(pf)
}

// SetStdoutFile sets a sys.stdout to a specified file descriptor.
func (py *Interpreter) SetStdoutFile(f *os.File) error {
	pf := py.fromFile(f, "w")
	defer pf.decRef()
	return py.SetStdoutObject(pf)
}

// SetStderrFile sets a sys.stderr to a specified file descriptor.
func (py *Interpreter) SetStderrFile(f *os.File) error {
	pf := py.fromFile(f, "w")
	defer pf.decRef()
	return py.SetStderrObject(pf)
}
