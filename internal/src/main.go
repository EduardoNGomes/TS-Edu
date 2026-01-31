package src

import (
	"fmt"
	"os"

	"gitbhub.com/eduardongomes/ts-edu/internal/src/domain"
	"gitbhub.com/eduardongomes/ts-edu/internal/src/env"
	"gitbhub.com/eduardongomes/ts-edu/internal/src/types"
)

func CreateSRCDIR() error {
	dir := "src"

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory src: %w", err)
	}

	if err := env.CreateEnvConfig(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/env: %w", err)
	}

	if err := types.CreateTypesDIR(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/@types: %w", err)
	}

	if err := domain.CreateDomainDIR(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/domain: %w", err)
	}
	return nil
}
