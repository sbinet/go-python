package python

type Object struct {
	ptr ObjectPtr
}

func (obj *Object) decRef() {
	if obj == nil || obj.ptr == nil {
		return
	}
	obj.ptr.DecRef()
}

func (obj *Object) setFinalizer() {
	// FIXME: set Go finalizer?
}

func newObject(ptr ObjectPtr) *Object {
	return &Object{ptr: ptr}
}
