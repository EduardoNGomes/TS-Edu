package src

import (
	"fmt"
	"os"

	"gitbhub.com/eduardongomes/ts-edu/internal/src/env"
)

func CreateSRCDIR() error {
	dir := "src"

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory src: %w", err)
	}

	if err := env.CreateEnvConfig(dir); err != nil {
		return fmt.Errorf("Failed to create dir env src: %w", err)
	}

	return nil
}
