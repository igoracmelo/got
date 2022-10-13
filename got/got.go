package got

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"sync"
)

type GitRepository struct {
	Dir string
}

type ExecResult struct {
	Stdout []byte
	Stderr []byte
	Error  error
}

type Got struct {
	Repositories []GitRepository
}

func (got *Got) ExecForRepoSingleMatch(pattern string, cmd *exec.Cmd) (*ExecResult, error) {
	matches := []int{}

	for i, repo := range got.Repositories {
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
	result := got.ExecForRepoIndex(index, *cmd)
	return result, nil
}

func (got *Got) ExecForRepoIndex(index int, cmd exec.Cmd) *ExecResult {
	if index >= len(got.Repositories) {
		return nil // TODO:
	}

	repo := got.Repositories[index]

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

func (got *Got) ExecForAll(cmd exec.Cmd) chan ExecResult {
	results := make(chan ExecResult)

	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(len(got.Repositories))

		for i := range got.Repositories {
			i := i
			go func() {
				results <- *got.ExecForRepoIndex(i, cmd)
				wg.Done()
			}()
		}

		wg.Wait()
		close(results)
	}()

	return results
}
