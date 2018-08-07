package python_test

import (
	"testing"

	python "github.com/sbinet/go-python"
)

func TestProgramName(t *testing.T) {
	const want = "foo.exe"
	python.Py_SetProgramName(want)
	name := python.Py_GetProgramName()
	if name != want {
		t.Fatalf("got=%q. want=%q", name, want)
	}
}

func TestPythonHome(t *testing.T) {
	const want = "/usr/lib/go-python"
	python.Py_SetPythonHome(want)
	got := python.Py_GetPythonHome()
	if got != want {
		t.Fatalf("got=%q. want=%q", got, want)
	}
}
