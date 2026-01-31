package http

import (
	"fmt"
	"os"
)

func CreateHTTP(dir string) error {
	httpDIR := dir + "/" + "http"

	if err := os.Mkdir(httpDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory http: %w", err)
	}

	if err := createControllers(httpDIR); err != nil {
		return fmt.Errorf("CONTROLLERS -> %w", err)
	}

	if err := createMiddlewares(httpDIR); err != nil {
		return fmt.Errorf("MIDDLEWARES -> %w", err)
	}

	filename := "index.ts"
	filepath := httpDIR + "/" + filename

	fileContent := `
	import { HCRoutes } from "./controllers/hc/routes";

export const routes = [HCRoutes];
`
	if err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm); err != nil {
		return fmt.Errorf("ERR on create %s : %w", filename, err)
	}

	return nil
}
