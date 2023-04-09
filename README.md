# Go-Todo

# Installation

1. install go from this link (https://go.dev/dl/)

# Setup for golang application

## enable dependency tracking using go mod init command

this creates a go.mod file that does dependency tracking
go mod init example/hello
go mod tidy -> authenticates the module, find the module and add it to the go.sum file for the dependency

## install dependencies

go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql

## get started!

# Setup for the todo api

docker-compose build
docker-compose up -d
go to the link localhost:3000/todos

# curl

## create

curl -X POST \
 -H "Content-Type: application/json" \
 -d '{"title": "", "description": ""}' \
 http://<link>/todo/create

Example:
curl -X POST \
 -H "Content-Type: application/json" \
 -d '{"title": "title", "description": "description"}' \
 http://localhost:3000/todo/create

## update

curl -X PUT \
 -H "Content-Type: application/json" \
 -d '{"title": "", "description": ""}' \
 http://<link>/update/{id}

Example:
curl -X PUT \
 -H "Content-Type: application/json" \
 -d '{"title": "bbb", "description": "bbb"}' \
 http://localhost:3000/update/4

## delete

curl -X DELETE \
 http://<link>/delete/{id}

Example:
curl -X DELETE \
 http://localhost:3000/delete/4
