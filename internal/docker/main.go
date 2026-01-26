package docker

import "fmt"

func DockerSetup() error {
	if err := createComposer(); err != nil {
		return fmt.Errorf("Error on create docker-compose.yml: Err => %w", err)
	}

	if err := createDockerfile(); err != nil {
		return fmt.Errorf("Error on create Dockerfile: Err => %w", err)
	}

	if err := createDockerIgnore(); err != nil {
		return fmt.Errorf("Error on create DockerIgnore: Err => %w", err)
	}

	return nil
}
