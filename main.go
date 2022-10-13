package main

import (
	"bytes"
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
	Stdout string
	Stderr string
	Error  error
}

func (g *Got) ExecForRepoSingleMatch(pattern string, cmd *exec.Cmd) {
	matches := []int{}

	for i, repo := range g.Repositories {
		if strings.Contains(repo.Dir, pattern) {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return // TODO: no matches
	}

	if len(matches) > 1 {
		return // TODO: many matches
	}

	index := matches[0]
	g.ExecForRepoIndex(index, cmd)
}

func (g *Got) ExecForRepoIndex(index int, cmd *exec.Cmd) *ExecResult {
	if len(g.Repositories) <= index {
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
		Stdout: stdoutBuf.String(),
		Stderr: stderrBuf.String(),
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
