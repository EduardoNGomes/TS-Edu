package dependencies

import "fmt"

func InstallCoreDependencies() error {
	if err := initPackageJSON(); err != nil {
		return fmt.Errorf("Failed to install dev dependencies: %w", err)
	}

	if err := installDependencies(); err != nil {
		return fmt.Errorf("Failed to install dev dependencies: %w", err)
	}

	if err := installDevDependencies(); err != nil {
		return fmt.Errorf("Failed to install dev dependencies: %w", err)
	}

	return nil

}
