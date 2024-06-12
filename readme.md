# BeliMang Backend Service

This is a backend service that provides api for cats matcher online

## Getting Started

### Prerequisites

- Go 1.22
- Postgres
- Golang Migrate CLI
- Docker

### Installation

1. Migrate the database schema

```sh
migrate -database "postgres://postgres:password@localhost:5432/eniqilodb?sslmode=disable" -path ./db/migrations -verbose up
```

2. Run the application

```sh
go run main.go
```

Specification App
https://openidea-projectsprint.notion.site/Cats-Social-9e7639a6a68748c38c67f81d9ab3c769
