# Go-Todo App

This is a simple Todo app that is used for the workshop hands-on session.
Users can create, update, and delete tasks.
The app is built using Golang/Gin and MySQL for data persistence.

## Features

- Retrieve all the tasks
- Retrieve a specific task by its ID
- Create a task
- Update a task
- Delete a task

## Installation

1. Clone the repository

```bash
git clone https://github.com/your-username/todo-app.git
```

2. Install Go from the link below

```bash
https://go.dev/dl/
```

3. Install Docker

- download and install the version of Docker for your OS from here:

```bash
https://docs.docker.com/get-docker/
```

- Note that for <b>Windows</b>, you will need to download wsl beforehand as well. You can find instructions for that here:

```bash
https://learn.microsoft.com/en-us/windows/wsl/install
```

## Start the development server

1. Start the docker containers
  ```bash
  docker-compose build
  docker-compose up -d
  ```

2. Visit the link or access inside the api container
  ```bash
  docker exec -it api /bin/sh
  ```
  - link: localhost:3000/todos

3. Send a request (Once implemented)

- create:
  ```bash
  curl -X POST \
   -H "Content-Type: application/json" \
   -d '{"title": "", "description": ""}' \
   http://<link>/todo/create

  Example:
  curl -X POST \
   -H "Content-Type: application/json" \
   -d '{"title": "title", "description": "description"}' \
   http://localhost:3000/todo/create
  ```

- update:
  ```bash
  curl -X PUT \
   -H "Content-Type: application/json" \
   -d '{"title": "", "description": ""}' \
   http://<link>/todo/update/{id}

  Example:
  curl -X PUT \
   -H "Content-Type: application/json" \
   -d '{"title": "bbb", "description": "bbb"}' \
   http://localhost:3000/todo/update/4
  ```

- delete:
  ```bash
  curl -X DELETE \
   http://<link>/todo/delete/{id}

  Example:
  curl -X DELETE \
   http://localhost:3000/todo/delete/4
  ```
