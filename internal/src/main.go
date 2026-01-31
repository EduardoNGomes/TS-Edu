package src

import (
	"fmt"
	"os"

	"gitbhub.com/eduardongomes/ts-edu/internal/src/domain"
	"gitbhub.com/eduardongomes/ts-edu/internal/src/env"
	"gitbhub.com/eduardongomes/ts-edu/internal/src/infra"
	"gitbhub.com/eduardongomes/ts-edu/internal/src/types"
)

func CreateSRCDIR() error {
	dir := "src"

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory src: %w", err)
	}

	if err := env.CreateEnvConfig(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/env: %w", err)
	}

	if err := types.CreateTypesDIR(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/@types: %w", err)
	}

	if err := domain.CreateDomainDIR(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/domain: %w", err)
	}

	if err := infra.CreateInfraDIR(dir); err != nil {
		return fmt.Errorf("Failed to create dir src/infra: %w", err)
	}

	filename := "app.ts"
	filepath := dir + "/" + filename

	fileContent := `import fastify from "fastify";
import fastifyJwt from "@fastify/jwt";
import { env } from "./env";
import fastifyCookie from "@fastify/cookie";
import multipart from "@fastify/multipart";
import { ZodError } from "zod";
import { routes } from "./infra/http";

export const app = fastify();

app.register(multipart, {
  limits: {
    fileSize: 10 * 1024 * 1024,
  },
});
app.register(fastifyJwt, {
  secret: env.JWT_SECRET,
  cookie: {
    cookieName: "refreshToken",
    signed: false,
  },
  sign: {
    expiresIn: "10m",
  },
});

app.register(fastifyCookie);

routes.forEach((route) => app.register(route));

app.setErrorHandler((error, _, reply) => {
  if (error instanceof ZodError) {
    return reply
      .status(400)
      .send({ message: "validation error.", issues: error.format() });
  }

  if (env.NODE_ENV !== "production") {
    console.error(error);
  } else {
    // TODO: Here we should log to an external tool like DataDog/NewRelic/Sentry
  }

  return reply.status(500).send({ message: "internal server error." });
});
`
	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	filename = "server.ts"
	filepath = dir + "/" + filename

	fileContent = `import { app } from "@/app";
import { env } from "./env";

app
  .listen({
    port: env.PORT,
  })
  .then(() => {
    console.log("server is lintening on PORT", env.PORT);
  });
`
	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
