# Fiber Clean Architecture With GORM

I offer this repository as a proposal for a clean architecture using fiber and GORM.
My inspirations were the gofiber recipe repositories.

### High-level Code Structure
The following structure is something that has evolved based on several projects from the gofiber recipes and go templates found on github.


```
- api --> You can create different output commands like Api rest, web, GRPC or any other technology.
  |- fiber --> structure fiber api
- cmd --> Main applications for this project.
- pkg --> Library code that's ok to use by external applications
  |- config
    |- config.go
- internal
  |- entity --> Application entities.
    |- user.go
    |- user_test.go
    |- ...
  |- infra --> I think the infrastructure should be separated by technology.
    |- gorm --> he gorm folder is responsible for managing the gorm infrastructure.
      |- database
        |- gorm.go
      |- repository
        |- user_gorm.go --> I think putting gorms tags on entities is not appropriate. Our application can simply stop using gorm and thus entities do not lose compatibility.
  |- usecase
    |- user
      |- interface.go
      |- service.go
```

## Endpoints

- GET   /             - Ping
- POST  /api/v1/users - Request: {"username", "password"} - Create a user
- GET   /api/v1/users - List all users
- GET   /api/v1/users/:id - List a user


## How to start

1. Run a postgresql container
```bash
docker run --name go-postgres -e POSTGRES_USER=golang -e POSTGRES_PASSWORD=12345678 -e POSTGRES_DB=fgcadb -p 5432:5432 -d postgres
```

2. Create a .env file
```bash
# Database settings:
DB_HOST="localhost"
DB_PORT=5432
DB_USER="golang"
DB_PASS="12345678"
DB_NAME="fgcadb"
DB_SSL_MODE="disable"
```

3. Run main.go
```bash
go run cmd/main.go
```

## ⚡️ Quick start

Alternative using docker-compose
```bash
docker-compose up
```

## 📦 Used packages

| Name                                                                  | Version   | Type       |
| --------------------------------------------------------------------- | --------- | ---------- |
| [go](https://go.dev/)                                                 | `v1.17`   | core       |
| [gofiber/fiber/v2](https://github.com/gofiber/fiber/v2)               | `v2.23.0` | core       |
| [asaskevich/govalidator](https://github.com/asaskevich/govalidator)   | `v0.0.2`  | core       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.4.0`  | config     |
| [gorm](https://gorm.io/gorm)                                          | `v1.22.4` | database   |
| [gorm/driver/postgres](https://gorm.io/driver/postgres)               | `v1.2.3`  | database   |
| [google/uuid](https://github.com/google/uuid)                         | `v1.3.0`  | utils      |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.0`  | test       |