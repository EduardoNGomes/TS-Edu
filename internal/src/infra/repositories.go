package infra

import (
	"fmt"
	"os"
)

func createRepositories(dir string) error {
	repositories := dir + "/" + "repositories"

	if err := os.Mkdir(repositories, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", repositories, err)
	}

	inMemoryDIR := repositories + "/" + "in-memory"

	if err := os.Mkdir(inMemoryDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", inMemoryDIR, err)
	}

	fileContent := `import { HCRepository } from "@/domain/repositories/hc";

export class InMemoryHCRepository implements HCRepository {
  async hello(name: string | null): Promise<string> {
    return "Hello " + name;
  }
}
`

	filename := "in-memory-hc-repository.ts"
	filepath := inMemoryDIR + "/" + filename

	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
