package python3

import (
	"github.com/sbinet/go-python/pytest"
	"testing"
)

func TestPython3(t *testing.T) {
	pytest.TestRuntime(t, Runtime)
}
