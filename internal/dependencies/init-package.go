package dependencies

import (
	"os"
	"os/exec"
)

func initPackageJSON() error {
	args := []string{
		"init",
		"--init-type=module",
	}

	cmd := exec.Command("pnpm", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil

}
