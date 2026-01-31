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

	return nil
}
