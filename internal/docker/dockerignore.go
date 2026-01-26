package docker

import (
	"os"
)

func createDockerIgnore() error {
	fileContent := `node_modules
.env
coverage
`

	err := os.WriteFile(".dockerignore", []byte(fileContent), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
