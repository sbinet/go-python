package python

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"gotest.tools/assert"
	"gotest.tools/golden"
)

func TestGoPython(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "-c", "print 1+1")
	cmd.Dir = "./cmd/go-python/"
	output, err := cmd.CombinedOutput()
	assert.NilError(t, err, string(output))
}

func TestCases(t *testing.T) {
	cases := []string{
		"cpickle",
		"errfetch",
		"modify-values",
		"issue61",
		"none-check",
	}
	os.Setenv("PYTHONPATH", os.ExpandEnv(".:$PYTHONPATH"))
	for _, dir := range cases {
		t.Run(golden.Path(dir), func(t *testing.T) {
			t.Parallel()
			cmd := exec.Command("go", "run", "main.go")
			cmd.Env = os.Environ()
			cmd.Dir = golden.Path(dir)
			got, err := cmd.CombinedOutput()
			assert.NilError(t, err, string(got))
			golden.Assert(t, string(got), filepath.Join(dir, "want.txt"))
		})
	}
}
