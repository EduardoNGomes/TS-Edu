package docker

import (
	"os"
)

func createDockerfile() error {
	fileContent := `FROM node:lts-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install --omit=dev


COPY . .

RUN npm run build

EXPOSE 8080

CMD ["npm", "run", "start"]
`

	err := os.WriteFile("Dockerfile", []byte(fileContent), os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
