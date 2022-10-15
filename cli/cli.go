package cli

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/igoracmelo/got/got"
)

// type GotParser struct{}

func ParseInstruction(cmd string, g *got.Got) *got.GotInstruction {
	r := regexp.MustCompile(`sh (.*?) (.*)`)

	matches := r.FindAllStringSubmatch(cmd, -1)[0]
	repoPattern := matches[1]
	subcommand := matches[2]

	fmt.Println(repoPattern)
	fmt.Println(subcommand)

	indexes := g.FindReposLike(repoPattern)

	return &got.GotInstruction{
		Command:      *exec.Command(subcommand),
		ReposIndexes: indexes,
	}
}

func ParseAndExec(cmd string, got *got.Got) error {
	return nil
}
