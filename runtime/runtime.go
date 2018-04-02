package runtime

import (
	"fmt"
	"os"
)

type Runtime interface {
	GetVersion() string
	GetCompiler() string

	IsInitialized() bool
	Initialize(signals bool)
	Finalize()

	Main(args []string) int

	EvalRuntime
	RunRuntime
	SysRuntime
	FileRuntime
	TypeRuntime
}

type Object interface {
	Valid() bool
	IsNone() bool
	DecRef()

	HasAttr(name Object) bool

	AsString() string
	StringCheck() bool
	StringCheckExact() bool
}

type DictObject interface {
	Object
}

type Frame interface {
	GetLineNumber() int
	GetFilename() Object

	Builtins() DictObject
	Globals() DictObject
	Locals() Object
}

type TraceType int

func (t TraceType) String() string {
	switch t {
	case TraceCall:
		return "Call"
	case TraceException:
		return "Exception"
	case TraceLine:
		return "Line"
	case TraceReturn:
		return "Return"
	case TraceCCall:
		return "C Call"
	case TraceCException:
		return "C Exception"
	case TraceCReturn:
		return "C Return"
	default:
		return fmt.Sprintf("TraceType(%d)", int(t))
	}
}

const (
	TraceCall = TraceType(iota + 1)
	TraceException
	TraceLine
	TraceReturn
	TraceCCall
	TraceCException
	TraceCReturn
)

type TraceFunc func(obj Object, frame Frame, what TraceType, arg Object) int

type EvalRuntime interface {
	EvalThreadsInitialized() bool
	EvalInitThreads()

	EvalSetTrace(fnc TraceFunc, obj Object)
	EvalGetFrame() Frame
}

type RunRuntime interface {
	RunString(command string) int
	RunFile(f *os.File) int
}

type SysRuntime interface {
	SysSetObject(name string, v Object) int
	SysGetObject(name string) Object
}

type FileRuntime interface {
	FromFile(f *os.File, mode string) Object
}

type TypeRuntime interface {
	None() Object
	True() Object
	False() Object
	FromString(v string) Object
	FromInt64(v int64) Object
	FromFloat64(v float64) Object
	FromBool(v bool) Object
}
