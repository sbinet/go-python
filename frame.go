package python

import (
	"fmt"
	"github.com/sbinet/go-python/runtime"
)

type Frame struct {
	ptr runtime.Frame
}

func (f *Frame) GetFilePos() *FileLine {
	if f == nil || f.ptr == nil {
		return nil
	}
	l := &FileLine{Line: f.ptr.GetLineNumber()}
	if name := f.ptr.GetFilename(); name.Valid() {
		l.Filename = name.AsString()
	}
	return l
}

type FileLine struct {
	Filename string
	Line     int
}

func (l FileLine) String() string {
	return fmt.Sprintf("%s:%d", l.Filename, l.Line)
}
