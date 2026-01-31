package dependencies

import (
	"os"
	"os/exec"
)

func installDevDependencies() error {
	libs := []string{
		"@types/bcryptjs",
		"@types/node",
		"@types/supertest",
		"@vitest/coverage-v8",
		"prisma",
		"supertest",
		"tsx",
		"typescript",
		"vite-tsconfig-paths",
		"vitest",
		"dotenv",
	}

	args := append([]string{"add"}, libs...)

	args = append(args, "-D")

	cmd := exec.Command("pnpm", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("pnpm", "add", "-D", "-E", "@biomejs/biome")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}
