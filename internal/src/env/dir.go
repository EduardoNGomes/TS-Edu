package env

import (
	"fmt"
	"os"
)

func createEnvDIR(dir string) error {
	envDIR := dir + "/" + "env"

	if err := os.Mkdir(envDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory env: %w", err)
	}

	filename := "index.ts"
	filepath := envDIR + "/" + filename
	fileContent := `import { z } from "zod";

const envSchema = z.object({
  JWT_SECRET:z.string(),
  NODE_ENV: z.enum(["dev", "production", "test"]).default("dev"),
  PORT: z.coerce.number().default(3333),
});

const _env = envSchema.safeParse(process.env);

if (_env.success === false) {
  console.log("Invalid environment variables", z.treeifyError(_env.error));
  throw new Error("Invalid environment variables");
}

export const env = _env.data;
`

	err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm)

	if err != nil {
		return fmt.Errorf("ERR on create env directory: %w", err)
	}

	return nil
}
