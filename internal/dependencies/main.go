package dependencies

import "fmt"

func InstallCoreDependencies() error {

	if err := installDevDependencies(); err != nil {
		return fmt.Errorf("Failed to install dev dependencies: %w", err)
	}

	return nil

}
