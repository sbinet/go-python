package python

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGoPython(t *testing.T) {
	cmd := exec.Command("go-python", "-c", "print 1+1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("go-python failed: %v", err)
	}
}

type pkg struct {
	path string
	want []byte
}

func testPkg(t *testing.T, table pkg) {
	workdir, err := ioutil.TempDir("", "go-python-")
	if err != nil {
		t.Fatalf("[%s]: could not create workdir: %v\n", table.path, err)
	}
	err = os.MkdirAll(workdir, 0644)
	if err != nil {
		t.Fatalf("[%s]: could not create workdir: %v\n", table.path, err)
	}
	defer os.RemoveAll(workdir)

	pypath := "." + string(os.PathListSeparator) + os.Getenv("PYTHONPATH")
	os.Setenv("PYTHONPATH", pypath)

	buf := new(bytes.Buffer)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin = os.Stdin
	cmd.Stdout = buf
	cmd.Stderr = buf
	cmd.Dir = table.path
	err = cmd.Run()
	if err != nil {
		t.Fatalf(
			"[%s]: error running go-python test: %v\n%v\n",
			table.path,
			err,
			string(buf.Bytes()),
		)
	}

	if !reflect.DeepEqual(string(buf.Bytes()), string(table.want)) {
		diffTxt := ""
		diffBin, diffErr := exec.LookPath("diff")
		if diffErr == nil {
			wantFile, wantErr := os.Create(filepath.Join(workdir, "want.txt"))
			if wantErr == nil {
				wantFile.Write(table.want)
				wantFile.Close()
			}
			gotFile, gotErr := os.Create(filepath.Join(workdir, "got.txt"))
			if gotErr == nil {
				gotFile.Write(buf.Bytes())
				gotFile.Close()
			}
			if gotErr == nil && wantErr == nil {
				cmd = exec.Command(diffBin, "-urN",
					wantFile.Name(),
					gotFile.Name(),
				)
				diff, _ := cmd.CombinedOutput()
				diffTxt = string(diff) + "\n"
			}
		}

		t.Fatalf("[%s]: error running go-python test:\nwant:\n%s\n\ngot:\n%s\n%s",
			table.path,
			string(table.want), string(buf.Bytes()),
			diffTxt,
		)
	}

}

func TestKwArgs(t *testing.T) {
	t.Parallel()
	testPkg(t, pkg{
		path: "tests/kw-args",
		want: []byte(`importing kwargs...
args=() kwds={}
args=() kwds={'a': 3}
`),
	})
}

func TestCPickle(t *testing.T) {
	t.Parallel()
	testPkg(t, pkg{
		path: "tests/cpickle",
		want: []byte(`hello [ foo ]
cPickle.dumps(foo) = "S'foo'\np1\n."
cPickle.loads("S'foo'\np1\n.") = "foo"
`),
	})
}

func TestErrFetch(t *testing.T) {
	t.Parallel()
	testPkg(t, pkg{
		path: "tests/errfetch",
		want: []byte("exc=<NULL>\nval=<NULL>\ntb=<NULL>\n"),
	})
}

func TestModifyValues(t *testing.T) {
	t.Parallel()
	testPkg(t, pkg{
		path: "tests/modify-values",
		want: []byte(`values.__name__: "values"
values.sval: "42"
values.ival: 666
sval='42'
ival=666
sval='42 is the answer'
ival=1666
`),
	})
}

func TestIssue61(t *testing.T) {
	t.Parallel()
	testPkg(t, pkg{
		path: "tests/issue61",
		want: []byte(`['i want this gone']
[]
`),
	})
}

func TestCheckNone(t *testing.T) {
	t.Parallel()
	testPkg(t, pkg{
		path: "tests/none-check",
		want: []byte(`type=<type 'NoneType'>, str=None, eq_none=true
`),
	})
}
