package dependencies

import (
	"os"
	"os/exec"
)

func installDependencies() error {
	libs := []string{
		"@fastify/cookie",
		"@fastify/jwt",
		"@prisma/client",
		"bcryptjs",
		"fastify",
		"fastify-multer",
		"tsdown",
		"zod",
	}

	args := append([]string{"install"}, libs...)

	cmd := exec.Command("pnpm", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}
