package docker

import (
	"os"
)

func createDockerfile() error {
	fileContent := `FROM node:lts-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install --omit=dev

COPY prisma ./prisma/

RUN npx prisma generate

COPY . .

RUN npm run build

EXPOSE 8080

CMD ["npm", "run", "start"]
`

	err := os.WriteFile("Dockerfile", []byte(fileContent), 0666)

	if err != nil {
		return err
	}

	return nil
}
