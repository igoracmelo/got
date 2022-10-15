package config

import (
	"os"
	"path"
	"strings"

	"github.com/igoracmelo/got/core"
)

func SaveConfig(configName string, g *core.Got) error {
	repoPaths := []string{}

	for _, repo := range g.Repositories {
		repoPaths = append(repoPaths, repo.Dir)
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	err = os.Mkdir(path.Join(dir, "got"), 0666)
	if err != nil {
		return err
	}

	data := strings.Join(repoPaths, "\n")
	configPath := path.Join(dir, "got", configName)

	err = os.WriteFile(configPath, []byte(data), 0666)
	return err
}

func LoadConfig(configName string, g *core.Got) error {
	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configPath := path.Join(dir, "got", configName)

	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	data := string(dataBytes)
	repoPaths := strings.Split(data, "\n")

	for _, repoPath := range repoPaths {
		g.Repositories = append(g.Repositories, core.GitRepository{
			Dir: repoPath,
		})
	}

	return nil
}
