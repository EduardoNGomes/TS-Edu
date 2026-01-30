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

	args := append([]string{"install"}, libs...)

	args = append(args, "-D")

	cmd := exec.Command("pnpm", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}
