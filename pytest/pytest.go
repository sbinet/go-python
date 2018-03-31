package pytest

import (
	"path/filepath"
	"testing"

	"bytes"
	"github.com/sbinet/go-python"
	"github.com/stretchr/testify/require"
	"time"
)

func TestRuntime(t *testing.T, r python.Runtime) {
	py := python.NewInterpreter(r)
	for _, c := range testTable {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := py.Initialize()
			require.NoError(t, err)
			defer py.Close()

			c.test(t, py)
		})
	}
}

const testDir = "../pytest/files/"

func pyFile(name string) string {
	return filepath.Join(testDir, name)
}

var testTable = []struct {
	name string
	test func(t *testing.T, py *python.Interpreter)
}{
	{name: "run string", test: testRunString},
	{name: "run file", test: testRunFile},
	{name: "main", test: testMain},
	{name: "exec out", test: testExec},
	{name: "exec error", test: testExecErr},
}

func testRunString(t *testing.T, py *python.Interpreter) {
	err := py.RunString("print('hi')")
	require.NoError(t, err)
}

func testRunFile(t *testing.T, py *python.Interpreter) {
	err := py.RunFile(pyFile("hi.py"))
	require.NoError(t, err)
}

func testMain(t *testing.T, py *python.Interpreter) {
	err := py.RunMain(pyFile("args.py"), "5")
	require.NoError(t, err)
}

func testExec(t *testing.T, py *python.Interpreter) {
	name := pyFile("args.py")
	cmd := py.Command(name, "5")
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	errc := make(chan error, 1)
	go func() {
		errc <- cmd.Run()
	}()
	select {
	case err := <-errc:
		require.NoError(t, err)
		require.Equal(t, buf.String(), "['"+name+"', '5']\n")
	case <-time.After(time.Second):
		require.Fail(t, "timeout")
	}
}

func testExecErr(t *testing.T, py *python.Interpreter) {
	cmd := py.Command(pyFile("raise.py"))
	errc := make(chan error, 1)
	go func() {
		errc <- cmd.Run()
	}()
	select {
	case err := <-errc:
		require.NotNil(t, err)
	case <-time.After(time.Second):
		require.Fail(t, "timeout")
	}
}
