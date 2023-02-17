package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/igoracmelo/got/cli"
	"github.com/igoracmelo/got/core"
)

func main() {
	g := &core.Got{
		Repositories: []core.GitRepository{
			{Dir: "/home/igor/Git/essia/essia-backend-auth"},
			{Dir: "/home/igor/Git/essia/essia-backend-catalog"},
			{Dir: "/home/igor/Git/essia/essia-backend-classwork"},
			{Dir: "/home/igor/Git/essia/essia-backend-ecommerce"},
			{Dir: "/home/igor/Git/essia/essia-backend-feature-flag"},
			{Dir: "/home/igor/Git/essia/essia-backend-media-optimizer"},
			{Dir: "/home/igor/Git/essia/essia-backend-pdf"},
			{Dir: "/home/igor/Git/essia/essia-backend-repository"},
			{Dir: "/home/igor/Git/essia/essia-backend-school-system"},
			{Dir: "/home/igor/Git/essia/essia-backend-trail"},
			{Dir: "/home/igor/Git/essia/essia-desktop"},
			{Dir: "/home/igor/Git/essia/essia-editor"},
			{Dir: "/home/igor/Git/essia/essia-frontend-backoffice"},
			{Dir: "/home/igor/Git/essia/essia-frontend-core"},
			{Dir: "/home/igor/Git/essia/essia-frontend-ecommerce"},
			{Dir: "/home/igor/Git/essia/essia-frontend-pdf-extractor"},
			{Dir: "/home/igor/Git/essia/essia-mobile"},
		},
	}

	instruction := cli.ArgsToInstruction(os.Args, g)
	results := g.RunInstruction(*instruction)

	for res := range results {
		// os.Stdout.Write(res.Stdout)
		fmt.Println(strings.Repeat("-", 100))
		// fmt.Println(strings.Repeat("-", 30))
		if res.Error != nil {
			fmt.Println("\033[1;31m[x] \033[1;37m" + path.Base(res.Repository.Dir) + "\033[0m")
		} else {
			fmt.Println("\033[1;32m[v] \033[1;37m" + path.Base(res.Repository.Dir) + "\033[0m")
		}

		stdout := string(res.Stdout)
		// stdoutLines := strings.Split(stdout, "\n")

		// if len(stdoutLines) > 10 {
		// 	stdout = strings.Join(stdoutLines[:10], "\n") + "\n(...)\n"
		// }

		stderr := string(res.Stderr)
		// stderrLines := strings.Split(stderr, "\n")

		// if len(stderrLines) > 10 {
		// 	stderr = strings.Join(stderrLines[:10], "\n") + "\n(...)\n"
		// }

		// stdout := strings.Join(strings.Split(string(res.Stdout), "\n")[:10], "\n")
		// stderr := strings.Join(strings.Split(string(res.Stderr), "\n")[:10], "\n")

		fmt.Print(stdout)
		fmt.Print(stderr)
		// os.Stdout.Write(res.Stdout)
		// os.Stderr.Write(res.Stderr)

		if len(res.Stdout) == 0 && len(res.Stderr) == 0 {
			fmt.Println("(empty)")
		}
		fmt.Println(strings.Repeat("-", 100))
		// fmt.Println()
		fmt.Println()
	}
}
