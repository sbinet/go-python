package python2

import (
	"github.com/sbinet/go-python/pytest"
	"testing"
)

func TestPython2(t *testing.T) {
	pytest.TestRuntime(t, Runtime)
}
