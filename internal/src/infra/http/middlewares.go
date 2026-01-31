package http

import (
	"fmt"
	"os"
)

func createMiddlewares(dir string) error {
	middlewaresDIR := dir + "/" + "middlewares"

	if err := os.Mkdir(middlewaresDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", middlewaresDIR, err)
	}

	filename := "verify-jwt.ts"
	filepath := middlewaresDIR + "/" + filename

	fileContent := `import { FastifyReply, FastifyRequest } from "fastify";

export async function verifyJwt(request: FastifyRequest, reply: FastifyReply) {
  try {
    await request.jwtVerify();
  } catch (error) {
    return reply.status(401).send({
      message: error,
    });
  }
}
`
	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
