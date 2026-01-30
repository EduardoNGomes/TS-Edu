package main

import (
	"fmt"

	"gitbhub.com/eduardongomes/ts-edu/internal/dependencies"
	"gitbhub.com/eduardongomes/ts-edu/internal/docker"
	"gitbhub.com/eduardongomes/ts-edu/internal/github"
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
}
