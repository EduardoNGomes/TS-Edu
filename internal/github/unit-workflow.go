package github

import "os"

func createUnitWorkflow(path string) error {
	filename := "run-unit-tests.yml"
	filepath := path + "/" + filename
	fileContent := `name: Run Unit Tests

on: [push]

jobs:
  run-unit-tests:
    name: Run Unit Tests
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-node@v3
        with:
          node-version: 24
          cache: 'npm'

      - run: npm ci

      - run: npm run test
	`

	err := os.WriteFile(filepath, []byte(fileContent), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
