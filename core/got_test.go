package core

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/igoracmelo/got/testutil"
	"github.com/stretchr/testify/assert"
)

func TestExecForRepoIndex_SingleRepository(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{dir},
		},
	}

	result := g.ExecForRepoIndex(0, *exec.Command("git", "status", "-s"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestExecForRepoIndex_MultiRepos(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{"should/ignore/this"},
			{"should/ignore/that"},
			{dir},
			{"and/ignore/this"},
		},
	}

	result := g.ExecForRepoIndex(2, *exec.Command("git", "status", "-s"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestExecForRepoIndex_MultiRepos_ShouldCommit(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{"should/ignore/this"},
			{"should/ignore/that"},
			{"and/ignore/this"},
			{dir},
		},
	}

	os.WriteFile(path.Join(dir, "file.txt"), []byte("hello world"), 0666)
	result := g.ExecForRepoIndex(3, *exec.Command("git", "add", "-A"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)

	result = g.ExecForRepoIndex(3, *exec.Command("git", "commit", "-m", "hello world"))
	assert.NoError(t, result.Error)
	assert.Empty(t, result.Stderr)

	result = g.ExecForRepoIndex(3, *exec.Command("git", "show", "--oneline"))
	assert.NoError(t, result.Error)
	assert.Contains(t, string(result.Stdout), "hello world")
	assert.Empty(t, result.Stderr)

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestExecForSingleMatch(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{"abcd"},
			{"def"},
			{dir},
			{"123"},
		},
	}

	// no matches
	_, err := g.ExecForRepoSingleMatch("xyz", exec.Cmd{})
	assert.Error(t, err)

	// multiple matches
	_, err = g.ExecForRepoSingleMatch("d", exec.Cmd{})
	assert.Error(t, err)

	// single match
	result, err := g.ExecForRepoSingleMatch("test", *exec.Command("ls", "-a"))
	assert.NoError(t, err)
	assert.NoError(t, result.Error)
	assert.Contains(t, string(result.Stdout), ".git")

	err = os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestExecForAll(t *testing.T) {
	dir1 := testutil.InitRandomRepo(t)
	dir2 := testutil.InitRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{dir1},
			{dir2},
		},
	}

	var err error

	results := g.ExecForAll(*exec.Command("ls", "-a"))

	for res := range results {
		assert.NoError(t, res.Error)
		assert.Contains(t, string(res.Stdout), ".git")
	}

	err = os.RemoveAll(dir1)
	assert.NoError(t, err)

	err = os.RemoveAll(dir2)
	assert.NoError(t, err)
}

func TestRunInstruction_SingleRepoInstruction(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &Got{
		Repositories: []GitRepository{
			{"should/ignore/that"},
			{"should/ignore/this"},
			{dir},
			{"and/ignore/this"},
		},
	}

	instruction := GotInstruction{
		Command:      *exec.Command("ls", "-a"),
		ReposIndexes: []int{2},
	}

	result := <-g.RunInstruction(instruction)
	assert.NoError(t, result.Error)
	assert.Contains(t, string(result.Stdout), ".git")
}
