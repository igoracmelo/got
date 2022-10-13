package main

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initRandomRepo(t *testing.T) string {
	dir, err := os.MkdirTemp("", "test-got-")
	assert.NoError(t, err)

	cmd := exec.Command("git", "init")
	cmd.Dir = dir

	outBytes, err := cmd.CombinedOutput()
	assert.NoError(t, err)
	assert.Contains(t, string(outBytes), "Initialized empty")

	return dir
}

func TestExecForRepoIndex_SingleRepository(t *testing.T) {
	dir := initRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{dir},
		},
	}

	result := g.ExecForRepoIndex(0, exec.Command("git", "status", "-s"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestExecForRepoIndex_MultiRepos(t *testing.T) {
	dir := initRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{"should/ignore/this"},
			{"should/ignore/that"},
			{dir},
			{"and/ignore/this"},
		},
	}

	result := g.ExecForRepoIndex(2, exec.Command("git", "status", "-s"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestExecForRepoIndex_MultiRepos_ShouldCommit(t *testing.T) {
	dir := initRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{"should/ignore/this"},
			{"should/ignore/that"},
			{"and/ignore/this"},
			{dir},
		},
	}

	os.WriteFile(path.Join(dir, "file.txt"), []byte("hello world"), 0666)
	result := g.ExecForRepoIndex(3, exec.Command("git", "add", "-A"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)

	result = g.ExecForRepoIndex(3, exec.Command("git", "commit", "-m", "hello world"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)

	result = g.ExecForRepoIndex(3, exec.Command("git", "show", "--oneline"))
	assert.NoError(t, result.Error)
	assert.Contains(t, result.Stdout, "hello world")
	assert.Empty(t, result.Stderr)

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}
