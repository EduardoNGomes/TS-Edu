# TS-Edu

TS-Edu is a command-line tool written in Go designed to bootstrap robust, highly structured, and production-ready Node.js applications using TypeScript. It automates the setup of a modern development environment, adhering to industry best practices and Clean Architecture principles.

## Motivation and Goals

The primary goal of TS-Edu is to eliminate the repetitive tasks involved in setting up a new Node.js project. By providing a pre-configured boilerplate, it ensures that developers can focus on implementing business logic rather than configuring build tools, linting rules, or directory structures. 

TS-Edu aims to:
- Standardize project structure across teams.
- Promote the use of Clean Architecture.
- Integrate modern tooling for performance and developer experience.
- Provide a CI/CD-ready foundation.

## Features

- **Fastify Framework:** High-performance and low-overhead web framework.
- **TypeScript First:** Fully typed development experience with optimized build processes.
- **Validation:** Schema validation and type inference using Zod.
- **Testing Setup:** Unit and E2E testing configuration with Vitest.
- **Linting and Formatting:** Fast and efficient code quality checks with Biome.
- **Containerization:** Ready-to-use Dockerfile and Docker Compose configuration.
- **CI/CD Integration:** Pre-configured GitHub Actions workflows for testing and linting.
- **Environment Management:** Type-safe environment variable handling.
- **Clean Architecture:** Domain-driven directory structure for better scalability and maintainability.

## Project Structure Overview

The generated project follows a structure inspired by Clean Architecture:

```text
src/
├── domain/           # Enterprise business rules (Entities, Repositories interfaces)
├── infra/            # External concerns (Frameworks, Database implementations, HTTP)
│   ├── http/         # Controllers, Routes, and Middlewares
│   ├── repositories/ # Concrete repository implementations
│   └── factories/    # Dependency injection factories
├── env/              # Environment variable configuration and validation
├── types/            # Global and third-party type definitions
├── app.ts            # Application initialization and plugin registration
└── server.ts         # Entry point for the HTTP server
```

## Requirements

To run the TS-Edu generator and the generated boilerplate, you will need:
- **Go** (for the generator tool)
- **Node.js** (LTS version recommended)
- **pnpm** (preferred and used by the generator)
- **Docker** and **Docker Compose** (optional, for containerized environments)

## Getting started

### Installation

#### Option 1: Download Pre-built Binary (Recommended)

You can download the latest version of **TS-Edu** directly from the [GitHub Releases](https://github.com/eduardongomes/ts-edu/releases) page. We provide pre-built binaries for Linux, macOS, and Windows.

1. Download the binary for your OS/Architecture (e.g., `ts-edu-linux-amd64`).
2. Make it executable (Linux/macOS).
3. Move it to a folder in your system's `$PATH`.

**Linux / macOS:**

```bash
# Example for Linux AMD64
wget https://github.com/eduardongomes/ts-edu/releases/latest/download/ts-edu-linux-amd64
chmod +x ts-edu-linux-amd64
sudo mv ts-edu-linux-amd64 /usr/local/bin/ts-edu

# Verify installation
ts-edu --help
```

**macOS users may need to approve the binary:**

If macOS blocks execution, run:
```bash
sudo xattr -dr com.apple.quarantine /usr/local/bin/ts-edu
```

**Windows:**

1. Download `ts-edu-windows-amd64.exe`.
2. Rename it to `ts-edu.exe`.
3. Move it to a folder that is in your system's `PATH` (e.g., `C:\Windows\System32` or a custom tools folder).

#### Option 2: Build from Source

If you have Go installed and prefer to build from the source code:

```bash
git clone https://github.com/eduardongomes/ts-edu.git
cd ts-edu
go build -o ts-edu ./cmd
```

Then move the binary to your PATH (e.g., `sudo mv ts-edu /usr/local/bin/`).

### Usage

Run the generator in your desired project directory:

```bash
./ts-edu
```

This will generate the standard directory structure, configuration files (`package.json`, `tsconfig.json`, `biome.json`, `vitest.config.ts`), and a base implementation of a Fastify server.

## Environment Variables

The generated project uses a `.env` file for configuration. A sample structure is provided:

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Port the server will listen on | 3333 |
| NODE_ENV | Environment mode (dev, test, production) | dev |
| JWT_SECRET | Secret key for JWT signing | - |

## Running in Development

After generating the project and installing dependencies, start the development server with hot-reload:

```bash
pnpm install
pnpm dev
```

The server uses `tsx` to run TypeScript files directly during development.

## Running Tests

TS-Edu scaffolds a testing environment using Vitest.

- **Run all tests:** `pnpm test`
- **Watch mode:** `pnpm test:watch`
- **Coverage report:** `pnpm test:coverage`

## Conventions and Best Practices Adopted

- **Clean Architecture:** Separation of concerns between business logic and infrastructure.
- **Dependency Inversion:** Use of interfaces for repositories to allow easy swapping of database providers.
- **Schema Validation:** All incoming data and environment variables are validated at runtime using Zod.
- **Functional Error Handling:** Consistent error responses and centralized error management.
- **Modern Tooling:** Preferring Biome over ESLint/Prettier for speed and Vitest over Jest for better TypeScript integration.

## How to Extend or Customize

The boilerplate is designed to be a starting point. To extend it:
1. **Adding Domain Logic:** Create new entities and repository interfaces in `src/domain`.
2. **Implementing Infrastructure:** Add concrete implementations in `src/infra/repositories`.
3. **Defining Routes:** Add new controllers and register routes in `src/infra/http`.
4. **Modifying Config:** Update `biome.json` or `tsconfig.json` to match your team's specific requirements.

## Contribution Guidelines

Contributions are welcome. Please follow these steps:
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes following the conventional commits standard.
4. Push to the branch.
5. Open a Pull Request.

Ensure that all Go code is formatted and tested before submitting.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
