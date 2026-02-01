package skill

import (
	"fmt"
	"os"
)

func CreateSkill() error {

	fileContent := "# Skill Implementation Guide\n\n" +
		"This document outlines the standard procedure for adding a new \"skill\" (feature/module) to the application. " +
		"A skill typically encapsulates a specific domain capability, spanning from the HTTP layer down to the domain logic and data persistence.\n\n" +

		"## 1. Overview\n\n" +
		"A **skill** represents a vertical slice of functionality. It is not a monolith; it is distributed across the Clean Architecture layers defined in the project.\n\n" +

		"### High-level Flow\n" +
		"1.  **HTTP Request** hits a route defined in `infra`.\n" +
		"2.  **Controller** receives the request.\n" +
		"3.  **Zod** validates input (Body/Query/Params).\n" +
		"4.  **Factory** assembles the Domain Service with its dependencies (Repositories).\n" +
		"5.  **Domain Service** executes business logic.\n" +
		"6.  **Repository** (Interface in Domain, Implementation in Infra) accesses data.\n" +
		"7.  **Result** is returned to Controller.\n" +
		"8.  **Controller** formats response or maps **Domain Errors** to HTTP Status Codes.\n\n" +

		"## 2. Directory Structure\n\n" +
		"You must distribute your files across `src/domain` and `src/infra`. Do not group everything in a single folder.\n\n" +

		"### Domain Layer (`src/domain`)\n" +
		"Pure TypeScript logic. No frameworks, no Zod, no HTTP concepts.\n\n" +

		"```text\n" +
		"src/domain/\n" +
		"├── errors/\n" +
		"│   └── <your-error>.ts       # Custom domain errors\n" +
		"├── repositories/\n" +
		"│   └── <skill>.ts            # Repository Interfaces\n" +
		"└── services/\n" +
		"    └── <skill>/\n" +
		"        ├── <skill>.ts        # Service Class\n" +
		"        └── <skill>.spec.ts   # Service Unit Tests\n" +
		"```\n\n" +

		"### Infrastructure Layer (`src/infra`)\n" +
		"Framework implementations (Fastify, Zod, Database).\n\n" +

		"```text\n" +
		"src/infra/\n" +
		"├── @factories/\n" +
		"│   └── <strategy>/<skill>/\n" +
		"│       └── <skill>.ts        # Factory to compose Service + Repo\n" +
		"├── http/\n" +
		"│   └── controllers/\n" +
		"│       └── <skill>/\n" +
		"│           ├── <action>.ts       # Controller (Handler)\n" +
		"│           ├── <action>.spec.ts  # Controller E2E/Integration Tests\n" +
		"│           └── routes.ts         # Route definitions\n" +
		"└── repositories/\n" +
		"    └── <strategy>/               # e.g., in-memory, prisma\n" +
		"        └── in-memory-<skill>-repository.ts\n" +
		"```\n\n" +

		"## 3. HTTP Validation with Zod\n\n" +
		"Validation happens **inside the controller**.\n\n" +

		"- **Location:** Define schemas directly within the controller function.\n" +
		"- **Parsing:** Use `.parse()`.\n" +
		"- **Error Handling:** `ZodError` is automatically caught by the global error handler in `app.ts` and converted to a **400 Bad Request**.\n\n" +

		"### Example (`src/infra/http/controllers/user/register.ts`)\n\n" +

		"```typescript\n" +
		"import type { FastifyReply, FastifyRequest } from \"fastify\";\n" +
		"import z from \"zod\";\n" +
		"import { makeRegisterUserService } from \"@/infra/@factories/in-memory/user/register\";\n\n" +
		"export async function register(request: FastifyRequest, reply: FastifyReply) {\n" +
		"    const registerService = makeRegisterUserService();\n\n" +
		"    const bodySchema = z.object({\n" +
		"        email: z.string().email(),\n" +
		"        password: z.string().min(6)\n" +
		"    });\n\n" +
		"    const { email, password } = bodySchema.parse(request.body);\n\n" +
		"    try {\n" +
		"        const result = await registerService.execute({ email, password });\n" +
		"        return reply.status(201).send(result);\n" +
		"    } catch (error) {\n" +
		"        return reply.status(500).send(error);\n" +
		"    }\n" +
		"}\n" +
		"```\n\n" +

		"## 4. Domain Layer\n\n" +
		"The domain layer describes *what* the software does, independent of *how* it is accessed.\n\n" +

		"### Rules\n" +
		"1.  **No Zod:** Use standard TypeScript interfaces for input/output.\n" +
		"2.  **No HTTP:** Never import `fastify` types here.\n" +
		"3.  **Dependency Inversion:** Services depend on Repository Interfaces, not implementations.\n\n" +

		"### Service Example (`src/domain/services/user/register.ts`)\n\n" +

		"```typescript\n" +
		"import type { UserRepository } from \"@/domain/repositories/user\";\n" +
		"import { UserAlreadyExistsError } from \"@/domain/errors/user-already-exists\";\n\n" +
		"interface RegisterUserRequest {\n" +
		"    email: string;\n" +
		"}\n\n" +
		"interface RegisterUserResponse {\n" +
		"    id: string;\n" +
		"}\n\n" +
		"export class RegisterUserService {\n" +
		"    constructor(private userRepository: UserRepository) {}\n\n" +
		"    async execute(input: RegisterUserRequest): Promise<RegisterUserResponse> {\n" +
		"        const exists = await this.userRepository.findByEmail(input.email);\n\n" +
		"        if (exists) {\n" +
		"            throw new UserAlreadyExistsError();\n" +
		"        }\n\n" +
		"        const user = await this.userRepository.create(input);\n\n" +
		"        return { id: user.id };\n" +
		"    }\n" +
		"}\n" +
		"```\n\n" +

		"### Domain Error Example (`src/domain/errors/user-already-exists.ts`)\n\n" +

		"```typescript\n" +
		"export class UserAlreadyExistsError extends Error {\n" +
		"    constructor() {\n" +
		"        super(\"User already exists.\");\n" +
		"        this.name = \"UserAlreadyExistsError\";\n" +
		"    }\n" +
		"}\n" +
		"```\n\n" +

		"## 7. Step-by-step: Creating a new skill\n\n" +
		"1.  **Domain:** Define the Repository Interface in `src/domain/repositories/`.\n" +
		"2.  **Domain:** Create custom Errors in `src/domain/errors/` (if needed).\n" +
		"3.  **Domain:** Implement the Service in `src/domain/services/<skill>/`.\n" +
		"4.  **Infra:** Implement the Repository in `src/infra/repositories/`.\n" +
		"5.  **Infra:** Create the Factory in `src/infra/@factories/`.\n" +
		"6.  **Infra:** Create the Controller in `src/infra/http/controllers/<skill>/`.\n" +
		"7.  **Infra:** Create `routes.ts` in the same folder.\n" +
		"8.  **Infra:** Register the new routes in `src/infra/http/index.ts`.\n"

	filename := "SKILLS.md"

	if err := os.WriteFile(filename, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("Error on create file: %s -> %w", filename, err)
	}

	return nil

}
