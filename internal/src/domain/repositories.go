package domain

import (
	"fmt"
	"os"
)

func createRepositories(dir string) error {
	filename := "hc.ts"
	repositoriesDIR := dir + "/" + "repositories"
	filepath := repositoriesDIR + "/" + filename

	if err := os.Mkdir(repositoriesDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", repositoriesDIR, err)
	}

	fileContent := `export interface HCRepository {
  hello(name: string | null): Promise<string>;
}
`

	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
