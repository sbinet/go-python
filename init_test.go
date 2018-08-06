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
