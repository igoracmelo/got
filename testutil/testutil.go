package testutil

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitRandomRepo(t *testing.T) string {
	dir, err := os.MkdirTemp("", "test-got-")
	assert.NoError(t, err)

	cmd := exec.Command("git", "init")
	cmd.Dir = dir

	outBytes, err := cmd.CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(outBytes), "Initialized empty")

	return dir
}
