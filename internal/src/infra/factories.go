package infra

import (
	"fmt"
	"os"
)

func createFactories(dir string) error {
	factoriesDIR := dir + "/" + "@factories"

	if err := os.Mkdir(factoriesDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", factoriesDIR, err)
	}

	inMemoryDIR := factoriesDIR + "/" + "in-memory"

	if err := os.Mkdir(inMemoryDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", inMemoryDIR, err)
	}

	hcDIR := inMemoryDIR + "/" + "hc"

	if err := os.Mkdir(hcDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", hcDIR, err)
	}

	fileContent := `import { HCService } from "@/domain/services/hc/hc";
import { InMemoryHCRepository } from "@/infra/repositories/in-memory/in-memory-hc-repository";

export function makeHCService() {
  const hcRepository = new InMemoryHCRepository();
  const hcService = new HCService(hcRepository);
  return hcService;
}
`

	filename := "hc.ts"
	filepath := hcDIR + "/" + filename

	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
