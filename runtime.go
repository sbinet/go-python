package python

import (
	"os"
)

type Runtime interface {
	IsInitialized() bool
	Initialize()
	Finalize()

	Main(args []string) int

	EvalRuntime
	RunRuntime
	SysRuntime
	FileRuntime
}

type ObjectPtr interface {
	Valid() bool
	DecRef()
}

type EvalRuntime interface {
	EvalThreadsInitialized() bool
	EvalInitThreads()
}

type RunRuntime interface {
	RunString(command string) int
	RunFile(f *os.File) int
}

type SysRuntime interface {
	SysSetObject(name string, v ObjectPtr) int
	SysGetObject(name string) ObjectPtr
}

type FileRuntime interface {
	FromFile(f *os.File, mode string) ObjectPtr
}
