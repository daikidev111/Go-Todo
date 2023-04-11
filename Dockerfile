FROM golang:1.20-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download && go mod verify
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN apk add --no-cache bash curl

EXPOSE 3000

ENTRYPOINT CompileDaemon -log-prefix=false --build="go build -o main /app/main.go" --command="./main"
