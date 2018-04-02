package pytest

import (
	"bytes"
	"log"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/sbinet/go-python"
	"github.com/sbinet/go-python/runtime"
	"github.com/stretchr/testify/require"
)

const runTrace = false

func TestRuntime(t *testing.T, r python.Runtime) {
	py := python.NewInterpreter(r)
	for _, c := range testTable {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := py.Initialize(false)
			require.NoError(t, err)
			defer py.Close()

			if runTrace {
				py.Trace(func(frame *python.Frame, what runtime.TraceType, arg runtime.Object) {
					log.Println(frame.GetFilePos(), what, arg)
				})
			}

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
	{name: "to string", test: testToString},
	{name: "run string", test: testRunString},
	{name: "run file", test: testRunFile},
	{name: "run file 2", test: testRunFile2},
	{name: "main", test: testMain},
	{name: "exec out", test: testExec},
	{name: "exec error", test: testExecErr},
}

func testToString(t *testing.T, py *python.Interpreter) {
	s := "Hello, 世界"
	got, ok := py.FromString(s).AsString()
	require.True(t, ok)
	require.Equal(t, s, got)
}

func testRunString(t *testing.T, py *python.Interpreter) {
	err := py.RunString("print('hi')")
	require.NoError(t, err)
}

func testRunFile(t *testing.T, py *python.Interpreter) {
	err := py.RunFile(pyFile("hi.py"))
	require.NoError(t, err)
}

func testRunFile2(t *testing.T, py *python.Interpreter) {
	err := py.RunFile(pyFile("hi8.py"))
	require.NoError(t, err)
}

func testMain(t *testing.T, py *python.Interpreter) {
	err := py.RunMain(pyFile("args.py"), "5")
	require.NoError(t, err)
}

func runCmd(t *testing.T, cmd *python.Cmd, py *python.Interpreter) error {
	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan struct{})
	go func() {
		defer wg.Done()
		select {
		case <-done:
		case <-time.After(time.Second):
			if f := py.GetFrame(); f != nil {
				log.Println(f.GetFilePos())
			}
			py.Close()
			require.Fail(t, "timeout")
		}
	}()
	err := cmd.Run() // keep on main thread
	close(done)
	wg.Wait()
	return err
}

func testExec(t *testing.T, py *python.Interpreter) {
	name := pyFile("args.py")
	cmd := py.Command(name, "5")
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := runCmd(t, cmd, py)
	require.NoError(t, err)
	require.Equal(t, buf.String(), "['"+name+"', '5']\n")
}

func testExecErr(t *testing.T, py *python.Interpreter) {
	cmd := py.Command(pyFile("raise.py"))

	err := runCmd(t, cmd, py)
	require.NotNil(t, err)
}
