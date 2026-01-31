package infra

import (
	"fmt"
	"os"
)

func CreateInfraDIR(dir string) error {
	infraDIR := dir + "/" + "infra"

	if err := os.Mkdir(infraDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory infra: %w", err)
	}

	if err := createFactories(infraDIR); err != nil {
		return fmt.Errorf("FACTORIES -> %w", err)
	}

	return nil
}
