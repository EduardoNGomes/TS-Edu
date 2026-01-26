package main

import (
	"fmt"

	"gitbhub.com/eduardongomes/ts-edu/internal/docker"
	"gitbhub.com/eduardongomes/ts-edu/internal/github"
)

func main() {
	if err := docker.DockerSetup(); err != nil {
		fmt.Printf("Error on DockerSetup: %v", err)
	}

	if err := github.CreateGitHubWorklow(); err != nil {
		fmt.Printf("Error on GitHubWorklflow: %v", err)
	}
}
