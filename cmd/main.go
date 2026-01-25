package main

import (
	"fmt"

	"gitbhub.com/eduardongomes/ts-edu/internal/docker"
)

func main() {
	if err := docker.DockerSetup(); err != nil {
		fmt.Printf("Error on DockerSetup: %v", err)
	}

}
