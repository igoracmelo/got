package cli

import (
	"os"
	"testing"

	"github.com/igoracmelo/got/got"
	"github.com/igoracmelo/got/testutil"
	"github.com/stretchr/testify/assert"
)

func TestParseInstruction(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &got.Got{}
	instruction := ParseInstruction("got sh repo ls", g)
	assert.Empty(t, instruction.ReposIndexes)
	assert.Equal(t, instruction.Command.Args, []string{"ls"})

	err := os.RemoveAll(dir)
	assert.NoError(t, err)
}

func TestParseAndExec(t *testing.T) {
	dir := testutil.InitRandomRepo(t)

	g := &got.Got{
		Repositories: []got.GitRepository{
			{"client"},
			{"server"},
			{dir},
		},
	}

	ParseAndExec("got sh test ls", g)
}
