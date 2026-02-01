package main

import (
	"fmt"
	"os"
	"os/exec"

	"gitbhub.com/eduardongomes/ts-edu/internal/configs"
	"gitbhub.com/eduardongomes/ts-edu/internal/dependencies"
	"gitbhub.com/eduardongomes/ts-edu/internal/docker"
	"gitbhub.com/eduardongomes/ts-edu/internal/github"
	"gitbhub.com/eduardongomes/ts-edu/internal/skill"
	"gitbhub.com/eduardongomes/ts-edu/internal/src"
)

func main() {
	if err := docker.DockerSetup(); err != nil {
		fmt.Printf("Error on DockerSetup: %v", err)
	}

	if err := github.CreateGitHubWorklow(); err != nil {
		fmt.Printf("Error on GitHubWorklflow: %v", err)

	}

	if err := dependencies.InstallCoreDependencies(); err != nil {
		fmt.Printf("Error on Install Deps: %v", err)
	}

	if err := src.CreateSRCDIR(); err != nil {
		fmt.Printf("Error on create SRC: %v", err)
	}

	if err := configs.CreateConfigFiles(); err != nil {
		fmt.Printf("Error on create Configs: %v", err)
	}

	if err := skill.CreateSkill(); err != nil {
		fmt.Printf("Error on create skill: %v", err)
	}

	cmd := exec.Command("pnpm", "lint")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error on apply lint")
	}

}
