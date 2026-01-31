package domain

import (
	"fmt"
	"os"
)

func createServices(dir string) error {
	servicesDIR := dir + "/" + "services"

	if err := os.Mkdir(servicesDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", servicesDIR, err)
	}

	hcDIR := servicesDIR + "/" + "hc"

	if err := os.Mkdir(hcDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", hcDIR, err)
	}

	fileContent := `import { HCRepository } from "@/domain/repositories/hc";

interface HCRequest {
  name: string | null;
}

interface HCResponse {
  result: string;
}

export class HCService {
  constructor(private hcRepository: HCRepository) {}

  async execute(input: HCRequest): Promise<HCResponse> {
    const result = await this.hcRepository.hello(input.name);

    return {
      result,
    };
  }
}
`

	filename := "hc.ts"
	filepath := hcDIR + "/" + filename
	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	fileContentTest := `import { HCService } from "./hc";
import { InMemoryHCRepository } from "@/infra/repositories/in-memory/in-memory-hc-repository";

let inMemoryHCRepository: InMemoryHCRepository;
let sut: HCService;

describe("Create HC Service", () => {
  beforeEach(async () => {
    inMemoryHCRepository = new InMemoryHCRepository();
    sut = new HCService(inMemoryHCRepository);
  });

  it("Should be able to create a new hc", async () => {
    const { result } = await sut.execute({ name: "World" });

    expect(result).toEqual("Hello World");
  });
});
`
	filenameTest := "hc.spec.ts"
	filepathTest := hcDIR + "/" + filenameTest
	if err := os.WriteFile(filepathTest, []byte(fileContentTest), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filenameTest, err)
	}

	return nil
}
