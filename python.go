package python

import (
	"fmt"
	"github.com/sbinet/go-python/runtime"
	"sync"
)

type Runtime = runtime.Runtime

func NewInterpreter(r Runtime) *Interpreter {
	return &Interpreter{r: r}
}

type Interpreter struct {
	mu sync.Mutex
	r  Runtime
}

// Initialize initializes the python interpreter and its GIL.
func (py *Interpreter) Initialize(signals bool) error {
	// make sure the python interpreter has been initialized
	if !py.r.IsInitialized() {
		py.r.Initialize(signals)
	}
	if !py.r.IsInitialized() {
		return fmt.Errorf("python: could not initialize the python interpreter")
	}

	// make sure the GIL is correctly initialized
	if !py.r.EvalThreadsInitialized() {
		py.r.EvalInitThreads()
	}
	if !py.r.EvalThreadsInitialized() {
		return fmt.Errorf("python: could not initialize the GIL")
	}
	return nil
}

// Close shutdowns the python interpreter by calling Finalize.
func (py *Interpreter) Close() error {
	py.r.Finalize()
	return nil
}
