package python

import (
	"fmt"
	"github.com/sbinet/go-python/runtime"
)

type Object struct {
	ptr runtime.Object
}

func (obj *Object) String() string {
	if obj == nil || obj.ptr == nil {
		return "<nil>"
	} else if !obj.ptr.Valid() {
		return "<invalid>"
	} else if obj.ptr.IsNone() {
		return "<None>"
	}
	if obj.ptr.StringCheck() {
		return obj.ptr.AsString()
	}
	return fmt.Sprint(obj.ptr)
}

func (obj *Object) AsString() (string, bool) {
	if obj.ptr.StringCheck() {
		return obj.ptr.AsString(), true
	}
	return "", false
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

func newObject(ptr runtime.Object) *Object {
	return &Object{ptr: ptr}
}
