package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type GitRepository struct {
	Dir string
}

type Got struct {
	Repositories []GitRepository
}

type ExecResult struct {
	Stdout []byte
	Stderr []byte
	Error  error
}

func (g *Got) ExecForRepoSingleMatch(pattern string, cmd *exec.Cmd) (*ExecResult, error) {
	matches := []int{}

	for i, repo := range g.Repositories {
		if strings.Contains(repo.Dir, pattern) {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return nil, errors.New("no matches found for pattern " + pattern) // TODO: no matches
	}

	if len(matches) > 1 {
		return nil, errors.New("multiple matches found for pattern " + pattern) // TODO: no matches
	}

	index := matches[0]
	result := g.ExecForRepoIndex(index, cmd)
	return result, nil
}

func (g *Got) ExecForRepoIndex(index int, cmd *exec.Cmd) *ExecResult {
	if index >= len(g.Repositories) {
		return nil // TODO:
	}

	repo := g.Repositories[index]

	cmd.Dir = repo.Dir
	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)

	cmd.Stdout = stdoutBuf
	cmd.Stderr = stderrBuf

	err := cmd.Run()

	return &ExecResult{
		Stdout: stdoutBuf.Bytes(),
		Stderr: stderrBuf.Bytes(),
		Error:  err,
	}
}

func main() {
	g := &Got{
		Repositories: []GitRepository{
			{Dir: "/home/igor/Git/pessoal/gitando"},
		},
	}

	g.ExecForRepoSingleMatch("gita", exec.Command("git", "log", "--oneline"))
}
