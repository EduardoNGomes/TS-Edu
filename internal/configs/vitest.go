package configs

import (
	"fmt"
	"os"
)

func createVitestConfigs() error {
	viteContent := `import { defineConfig } from "vitest/config";
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
  test: {
    globals: true,
    environment: "node",
    setupFiles: ["./vitest.setup.ts"],
  },
  plugins: [tsconfigPaths()],
});
`
	if err := os.WriteFile("vite.config.ts", []byte(viteContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create vite.config.ts: %w", err)
	}

	viteContentE2E := `import { defineConfig } from "vitest/config";
import tsConfigPaths from "vite-tsconfig-paths";

export default defineConfig({
  test: {
    include: ["**/*.e2e-spec.ts"],
    globals: true,
    root: "./",
    setupFiles: ["./vitest.setup-e2e.ts"],
  },
  plugins: [tsConfigPaths()],
});
`

	if err := os.WriteFile("vite.config-e2e.ts", []byte(viteContentE2E), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create vite.config-e2e.ts: %w", err)
	}

	vitestSetup := `import { config } from "dotenv";

config({ path: ".env.test" });
`

	if err := os.WriteFile("vitest.setup.ts", []byte(vitestSetup), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create vitest.setup.ts: %w", err)
	}

	vitestSetupE2E := `import { config } from "dotenv";

import { randomUUID } from "node:crypto";
import { envSchema } from "@/env";

config({ path: ".env", override: true });
config({ path: ".env.test", override: true });

const env = envSchema.parse(process.env);

function generateUniqueDataBaseURL(schemaId: string) {
  if (!env.DATABASE_URL) {
    throw new Error("DATABASE_URL is required");
  }
  const url = new URL(env.DATABASE_URL);

  url.searchParams.set("schema", schemaId);

  return url.toString();
}

const schemaId = randomUUID();
beforeAll(async () => {
  const databaseURL = generateUniqueDataBaseURL(schemaId);

  process.env.DATABASE_URL = databaseURL;

  /**
   * EXEC DB MIGRATIONS AND SEED
   * Here you can run your migrations and seeders
   *
   * Example:
   *
   * execSync("pnpm prisma migrate deploy");
   * execSync("pnpm prisma db seed");
   */
});

afterAll(async () => {
  /** Here you can close your database connection and clean up
   * Example:
   * await prisma.$executeRawUnsafe('DROP SCHEMA IF EXISTS "${schemaId}" CASCADE')
   * await prisma.$disconnect()
   */
});
`

	if err := os.WriteFile("vitest.setup-e2e.ts", []byte(vitestSetupE2E), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create vitest.setup-e2e.ts: %w", err)
	}

	return nil
}
