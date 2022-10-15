package cli

import (
	"os/exec"

	"github.com/igoracmelo/got/core"
)

func ArgsToInstruction(args []string, g *core.Got) *core.GotInstruction {
	cmdName := args[1]

	instruction := new(core.GotInstruction)
	gitArgs := []string{
		"-c", "color.status=always",
		"-c", "color.ui=always",
	}

	if cmdName == "as" {
		pattern := args[2]
		instruction.ReposIndexes = g.FindReposLike(pattern)
		gitArgs = append(gitArgs, args[3:]...)
		instruction.Command = *exec.Command("git", gitArgs...)
	} else {
		gitArgs = append(gitArgs, args[1:]...)
		instruction.Command = *exec.Command("git", gitArgs...)
		for i := range g.Repositories {
			instruction.ReposIndexes = append(instruction.ReposIndexes, i)
		}
	}

	return instruction
}
