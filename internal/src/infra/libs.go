package infra

import (
	"fmt"
	"os"
)

func createLibs(dir string) error {
	libsDIR := dir + "/" + "@libs"

	if err := os.Mkdir(libsDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", libsDIR, err)
	}

	fileContent := ""

	filename := ".gitkeep"
	filepath := libsDIR + "/" + filename

	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
