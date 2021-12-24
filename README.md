# Fiber Clean Architecture With GORM

I offer this repository as a proposal for a clean architecture using fiber and GORM.
My inspirations were the gofiber recipe repositories.

### High-level Code Structure
The following structure is something that has evolved based on several projects from the gofiber recipes and go templates found on github.


```
- cmd --> folder that has one or more application entry points.
  |- api --> You can create different output commands like Api rest, web, GRPC or any other technology.
    |- handler -->
    |- router -->
    |- app.go
- config --> It's a utility folder. I wonder if the .envs files could be in here.
  |- config.go
- entity --> Application entities.
  |- user.go
  |- user_test.go
  |- ...
- infra --> I think the infrastructure should be separated by technology.
  |- gorm --> he gorm folder is responsible for managing the gorm infrastructure.
    |- database
      |- gorm.go
    |- repository
      |- user_gorm.go --> I think putting gorms tags on entities is not appropriate. Our application can simply stop using gorm and thus entities do not lose compatibility.
- usecase
  |- user
    |- interface.go
    |- service.go
```

## Endpoints

- GET / - Ping
- POST /api/v1/users - {"username", "password"}

## ⚡️ Quick start

1. Run a postgresql container
```bash
docker run --name go-postgres -e POSTGRES_USER=golang -e POSTGRES_PASSWORD=12345678 -e POSTGRES_DB=fgcadb -p 5432:5432 -d postgres
```

2. Run main.go
```bash
go run main.go
```

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