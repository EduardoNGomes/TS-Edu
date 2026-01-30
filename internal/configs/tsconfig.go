package configs

import (
	"os"
)

func createTSConfig() error {
	fileContent := `{
  "compilerOptions": {
    "target": "es2020",
    "module": "commonjs",
    "baseUrl": "./",
    "paths": {"@/*":["./src/*"]}, 
    "forceConsistentCasingInFileNames": true,
    "strict": true,
    "skipLibCheck": true,
	"esModuleInterop": true,
	"types": [
      "vitest/globals"
    ]
  }
}
`
	err := os.WriteFile("tsconfig.json", []byte(fileContent), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
