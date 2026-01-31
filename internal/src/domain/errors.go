package domain

import (
	"fmt"
	"os"
)

func createErrors(dir string) error {
	filename := "unauthorized.ts"
	errorsDIR := dir + "/" + "errors"
	filepath := errorsDIR + "/" + filename

	if err := os.Mkdir(errorsDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory errors: %w", err)
	}

	fileContent := `export class UnauthorizedError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "UnauthorizedError";
  }
}
`

	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
