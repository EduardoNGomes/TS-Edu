package domain

import (
	"fmt"
	"os"
)

func CreateDomainDIR(dir string) error {
	domainDIR := dir + "/" + "domain"

	if err := os.Mkdir(domainDIR, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory domain: %w", err)
	}

	if err := createErrors(domainDIR); err != nil {
		return err
	}

	if err := createRepositories(domainDIR); err != nil {
		return err
	}

	if err := createServices(domainDIR); err != nil {
		return err
	}

	return nil
}
