FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download && go mod verify
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 3000

ENTRYPOINT CompileDaemon -log-prefix=false --build="go build -o main /app/main.go" --command="./main"
