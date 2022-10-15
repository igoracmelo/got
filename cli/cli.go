package cli

import (
	"os"
	"os/exec"
	"path"

	"github.com/igoracmelo/got/config"
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
	} else if cmdName == "register" {
		g.Repositories = []core.GitRepository{}
		configName := args[2]
		repoPaths := args[3:]
		workingDir, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		for _, repopath := range repoPaths {
			if !path.IsAbs(repopath) {
				repopath = path.Join(workingDir, repopath)
			}
			g.Repositories = append(g.Repositories, core.GitRepository{Dir: repopath})
		}

		config.SaveConfig(configName, g)
	} else {
		gitArgs = append(gitArgs, args[1:]...)
		instruction.Command = *exec.Command("git", gitArgs...)
		for i := range g.Repositories {
			instruction.ReposIndexes = append(instruction.ReposIndexes, i)
		}
	}

	return instruction
}
