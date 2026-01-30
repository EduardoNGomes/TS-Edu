package env

import (
	"fmt"
	"os"
)

func createEnvFile() error {
	fileContent := `DATABASE_URL="postgresql://docker:docker@localhost:5432/db?schema=public"

JWT_SECRET=secret
`

	if err := os.WriteFile(".env", []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create .env: %w", err)
	}

	if err := os.WriteFile(".env.test", []byte(fileContent), os.ModePerm); err != nil {

		return fmt.Errorf("ERR on create .env.test: %w", err)
	}
	return nil
}
