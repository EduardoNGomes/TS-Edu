package docker

import "fmt"

func DockerSetup() error {
	if err := createComposer(); err != nil {
		fmt.Printf("Error on create docker-compose.yml: Err => %v", err)
	}

	if err := createDockerfile(); err != nil {
		fmt.Printf("Error on create Dockerfile: Err => %v", err)
	}

	if err := createDockerIgnore(); err != nil {
		fmt.Printf("Error on create DockerIgnore: Err => %v", err)
	}

	return nil
}
