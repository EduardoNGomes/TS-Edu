package infra

import (
	"fmt"
	"os"

	"gitbhub.com/eduardongomes/ts-edu/internal/src/infra/http"
)

func CreateInfraDIR(dir string) error {
	infraDIR := dir + "/" + "infra"

	if err := os.Mkdir(infraDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory infra: %w", err)
	}

	if err := createFactories(infraDIR); err != nil {
		return fmt.Errorf("FACTORIES -> %w", err)
	}

	if err := createLibs(infraDIR); err != nil {
		return fmt.Errorf("LIBS -> %w", err)
	}

	if err := http.CreateHTTP(infraDIR); err != nil {
		return fmt.Errorf("HTTP -> %w", err)
	}

	return nil
}
