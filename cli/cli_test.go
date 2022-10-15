package cli

import (
	"os"
	"testing"

	"github.com/igoracmelo/got/core"
	"github.com/igoracmelo/got/testutil"
	"github.com/stretchr/testify/assert"
)

func TestArgsToInstruction(t *testing.T) {
	first := testutil.InitRandomRepoPrefix(t, "first")
	sec := testutil.InitRandomRepoPrefix(t, "sec")
	defer os.RemoveAll(first)
	defer os.RemoveAll(sec)

	g := &core.Got{
		Repositories: []core.GitRepository{
			{Dir: first},
			{Dir: sec},
		},
	}

	instruction := ArgsToInstruction([]string{"got", "as", "first", "status"}, g)
	assert.Equal(t, "git", instruction.Command.Args[0])
	assert.Contains(t, instruction.Command.Args, "status")
	assert.Len(t, instruction.ReposIndexes, 1)

	instruction = ArgsToInstruction([]string{"got", "pull"}, g)
	assert.Equal(t, "git", instruction.Command.Args[0])
	assert.Contains(t, instruction.Command.Args, "pull")
	assert.Len(t, instruction.ReposIndexes, 2)
}
