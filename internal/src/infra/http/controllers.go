package http

import (
	"fmt"
	"os"
)

func createControllers(dir string) error {
	controllersDIR := dir + "/" + "controllers"

	if err := os.Mkdir(controllersDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", controllersDIR, err)
	}

	hcDIR := controllersDIR + "/" + "hc"

	if err := os.Mkdir(hcDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", hcDIR, err)
	}

	filename := "hc.ts"
	filepath := hcDIR + "/" + filename

	hcContent := `import { makeHCService } from "@/infra/@factories/in-memory/hc/hc";
import { FastifyReply, FastifyRequest } from "fastify";
import z from "zod";

export async function hc(request: FastifyRequest, reply: FastifyReply) {
  const hcService = makeHCService();

  const paramSchema = z.object({
    name: z.string().optional(),
  });

  const { name } = paramSchema.parse(request.query);

  try {
    const r = await hcService.execute({ name: name || null });

    return reply.status(200).send({
      message: r,
    });
  } catch (error) {
    return reply.status(500).send(error);
  }
}
`
	if err := os.WriteFile(filepath, []byte(hcContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	filename = "hc.spec.ts"
	filepath = hcDIR + "/" + filename

	hcContentTest := `import { app } from "@/app";
	import request from "supertest";

	describe("HC(e2e)", async () => {
		beforeEach(async () => {
			await app.ready();
		});

		it("should return 200", async () => {
			const data = await request(app.server).get("/hc");

			expect(data.status).toEqual(200);
		});
	});
`

	if err := os.WriteFile(filepath, []byte(hcContentTest), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	filename = "hcSign.ts"
	filepath = hcDIR + "/" + filename

	hcSignContent := `import { makeHCService } from "@/infra/@factories/in-memory/hc/hc";
import { FastifyReply, FastifyRequest } from "fastify";
import z from "zod";

export async function hcSign(request: FastifyRequest, reply: FastifyReply) {
  const hcService = makeHCService();

  const paramSchema = z.object({
    name: z.string().optional(),
  });

  const { name } = paramSchema.parse(request.query);

  try {
    const r = await hcService.execute({ name: name || null });

    return reply.status(200).send({
      message: r,
    });
  } catch (error) {
    return reply.status(500).send(error);
  }
}
`

	if err := os.WriteFile(filepath, []byte(hcSignContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	filename = "hcSign.spec.ts"
	filepath = hcDIR + "/" + filename

	hcSignContentTest := `import { app } from "@/app";
import request from "supertest";

describe("HC Sign(e2e)", async () => {
  beforeEach(async () => {
    await app.ready();
  });

  it("should return a error because there isnt authorization", async () => {
    const data = await request(app.server).get("/hc/sign");

    expect(data.status).toEqual(401);
  });
});
`
	if err := os.WriteFile(filepath, []byte(hcSignContentTest), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	filename = "routes.ts"
	filepath = hcDIR + "/" + filename

	routesContent := `import { FastifyInstance } from "fastify";
import { hc } from "./hc";
import { hcSign } from "./hcSign";
import { verifyJwt } from "@/infra/http/middlewares/verify-jwt";

export async function HCRoutes(app: FastifyInstance) {
  app.get("/hc", hc);
  app.get("/hc/sign", { onRequest: [verifyJwt] }, hcSign);
}
`
	if err := os.WriteFile(filepath, []byte(routesContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
