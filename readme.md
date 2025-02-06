# APIs Services of 'Learn Anywhere' web application

### Tech Stack
- Go + Fiber
- PostgreSQL
- MongoDB

### Installation
```bash
cp .env.example .env
```

```bash
docker compose up -d
```

# APIs Services of 'Learn Anywhere' web application

### Tech Stack
- Go + Fiber
- PostgreSQL
- MongoDB

### Installation
```bash
cp .env.example .env
```

```bash
docker compose up -d
```

```
.
├── Dockerfile
├── configs
│   ├── config.go
│   └── db.go
├── docker-compose.yml
├── domain
│   ├── models
│   ├── repositories
│   ├── requests
│   ├── responses
│   └── usecases
├── go.mod
├── go.sum
├── infrastructure
│   └── auth
├── internal
│   └── adapters
│       ├── middlewares
│       │   └── middleware.go
│       ├── pg
│       ├── rest
│       │   └── routes
│       └── sqlx
├── main.go
├── readme.md
└── tmp
    ├── build-errors.log
    └── main
```
