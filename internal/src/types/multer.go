package types

import (
	"fmt"
	"os"
)

func createMulterType(dir string) error {
	filename := "fastify-multer.d.ts"
	filepath := dir + "/" + filename
	fileContent := `import { FastifyRequest, FastifyReply } from 'fastify'

declare module 'fastify' {
  interface FastifyRequest {
    file: {
      filename: string
      path: string
    }
  }
  interface FastifyReply {
    sendFile: (string) => void
  }
}
`

	err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm)

	if err != nil {
		return fmt.Errorf("ERR on create fastify-multer.d.ts : %w", err)
	}

	return nil
}
