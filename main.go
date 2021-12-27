package main

import (
	"github.com/vinigracindo/fiber-gorm-clean-architecture/cmd/api"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/config"
)

func main() {
	config.LoadEnvironmentFile(".env")
	api.Run(3000)
}
