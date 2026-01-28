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
		"eslint",
		"prisma",
		"supertest",
		"tsx",
		"typescript",
		"vite-tsconfig-paths",
		"vitest",
	}

	args := append([]string{"install"}, libs...)

	cmd := exec.Command("npm", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}
