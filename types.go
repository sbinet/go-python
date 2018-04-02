package python

func (py *Interpreter) FromString(s string) *Object {
	o := newObject(py.r.FromString(s))
	o.setFinalizer()
	return o
}
