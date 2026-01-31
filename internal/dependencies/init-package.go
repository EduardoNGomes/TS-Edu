package dependencies

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
)

func initPackageJSON() error {
	args := []string{
		"init",
		"--init-type=module",
	}

	cmd := exec.Command("pnpm", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("whoami")

	out, err := cmd.Output()

	if err != nil {
		return err
	}

	cmd = exec.Command("touch", ".env")

	if err := cmd.Run(); err != nil {
		return err
	}

	person := strings.TrimSpace(string(out))

	f, err := os.ReadFile("package.json")

	if err != nil {
		return err
	}

	var decodedJSON map[string]any

	if err := json.Unmarshal(f, &decodedJSON); err != nil {
		return err
	}

	for k := range decodedJSON {

		switch k {
		case "main":
			{
				decodedJSON[k] = "server.ts"
			}
		case "author":
			{
				decodedJSON[k] = person
			}
		case "scripts":
			{
				if scripts, ok := decodedJSON[k].(map[string]any); ok {
					scripts["dev"] = "tsx watch --env-file=.env src/server.ts"
					scripts["start"] = "node --env-file=.env build/server.js"
					scripts["build"] = "tsdown src"
					scripts["test"] = "vitest run"
					scripts["test:watch"] = "vitest"
					scripts["test:coverage"] = "vitest run --coverage"
					scripts["lint"] = "biome check --write --unsafe"
				}
			}
		default:
			{
				continue
			}
		}
	}

	encodedJSON, err := json.MarshalIndent(decodedJSON, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile("package.json", encodedJSON, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
