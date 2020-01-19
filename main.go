package main

import (
	"github.com/claudeseo/railgun/src"
	"github.com/claudeseo/railgun/src/config"
	"github.com/claudeseo/railgun/src/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.Init()
	database.Init()
	src.Init()
}
