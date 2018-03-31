package python

import (
	"os"
)

func (py *Interpreter) RunString(command string) error {
	py.mu.Lock()
	defer py.mu.Unlock()
	rcode := py.r.RunString(command)
	if rcode == 0 {
		return nil
	}
	return &RunError{Code: rcode}
}

func (py *Interpreter) RunFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	py.mu.Lock()
	defer py.mu.Unlock()
	rcode := py.r.RunFile(f)
	if rcode == 0 {
		return nil
	}
	return &RunError{Code: rcode, File: filename}
}

func (py *Interpreter) Main(args []string) error {
	py.mu.Lock()
	defer py.mu.Unlock()
	rcode := py.r.Main(args)
	if rcode == 0 {
		return nil
	}
	return &RunError{Code: rcode}
}

func (py *Interpreter) RunMain(name string, args ...string) error {
	return py.Main(append([]string{os.Args[0], name}, args...))
}
