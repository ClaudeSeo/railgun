package main

import (
	"strconv"

	"github.com/claudeseo/railgun/config"
	"github.com/claudeseo/railgun/database"
	"github.com/claudeseo/railgun/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.Init()
	database.Init()
	cfg := config.GetConfig()
	router := routes.Init()
	endpoint := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	router.Run(endpoint)
}
