package python

import (
	"os/exec"
	"testing"
)

func TestGoPython(t *testing.T) {
	cmd := exec.Command("go-python", "-c", "print 1+1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("go-python failed: %v", err)
	}
}

// EOF
