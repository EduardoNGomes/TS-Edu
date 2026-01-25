package docker

import (
	"os"
)

func CreateComposer() error {
	fileContent := `version: '3'

services:
  postgres_db:
    image: bitnami/postgresql
    ports:
      - 5432:5432
    environment:
      - POSTGRESQL_USERNAME=docker
      - POSTGRESQL_PASSWORD=docker
      - POSTGRESQL_DATABASE=db
	`

	err := os.WriteFile("docker-compose.yml", []byte(fileContent), 0666)

	if err != nil {
		return err
	}

	return nil
}
