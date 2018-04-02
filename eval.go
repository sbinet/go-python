package python

import "github.com/sbinet/go-python/runtime"

type TraceFunc func(frame *Frame, what runtime.TraceType, arg runtime.Object)

func (py *Interpreter) Trace(fnc TraceFunc) {
	var obj runtime.Object // TODO: unique value
	py.r.EvalSetTrace(func(_ runtime.Object, frame runtime.Frame, what runtime.TraceType, arg runtime.Object) int {
		fnc(&Frame{ptr: frame}, what, arg)
		return 0
	}, obj)
}

func (py *Interpreter) GetFrame() *Frame {
	f := py.r.EvalGetFrame()
	if f == nil {
		return nil
	}
	return &Frame{ptr: f}
}
