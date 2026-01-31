package types

import (
	"fmt"
	"os"
)

func CreateTypesDIR(dir string) error {
	typesDIR := dir + "/" + "@types"

	if err := os.Mkdir(typesDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory @types: %w", err)
	}

	if err := createJWTType(typesDIR); err != nil {
		return err
	}

	if err := createMulterType(typesDIR); err != nil {
		return err
	}

	return nil
}
