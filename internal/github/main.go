package github

import (
	"fmt"
	"os"
)

func CreateGitHubWorklow() error {
	githubDir := ".github"
	workFlowDir := "workflows"

	fullDir := githubDir + "/" + workFlowDir

	if err := os.Mkdir(githubDir, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory .github: %w", err)
	}

	if err := os.Mkdir(fullDir, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory workflows: %w", err)
	}

	if err := createE2EWorkflow(fullDir); err != nil {
		return fmt.Errorf("Failed to create E2E workflow: %w", err)
	}

	if err := createUnitWorkflow(fullDir); err != nil {
		return fmt.Errorf("Failed to create Unit workflow: %w", err)
	}

	return nil
}
