package main

import (
	"fmt"

	"gitbhub.com/eduardongomes/ts-edu/internal/docker"
)

func main() {
	if err := docker.CreateComposer(); err != nil {
		fmt.Printf("Error on create docker-compose.yml: Err => %v", err)
	}

}
