package types

import (
	"fmt"
	"os"
)

func createJWTType(dir string) error {
	filename := "fastify-jwt.d.ts"
	filepath := dir + "/" + filename
	fileContent := `import "@fastify/jwt";

declare module "@fastify/jwt" {
  export interface FastifyJWT {
    user: {
      sign: {
        sub: string;
      };
    };
  }
}
`

	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create fastify-jwt.d.ts : %w", err)
	}

	return nil
}
