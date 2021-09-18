This project is a simple REST API for user management built with Go.

## Features

- CRUD user
- Login and logout
- Access and refresh token
- Validation
- Authorization
- API documentation (Swagger)

## Quick Setup

Before you proceed, please make sure these programs are present:

- Go (1.16)
- MySQL (8.0)
- Golang Migrate

Run `go mod tidy` to install required packages.

Copy file `.env.example` to `.env` and update its content.

Migrate database by running `make migrate-up`.

## Running The Program

Run `go run main.go` to start the program. It should be available at http://localhost:8080.

For API documentation, visit http://localhost:8080/swagger/index.html.

## Demo

No demo link yet.
