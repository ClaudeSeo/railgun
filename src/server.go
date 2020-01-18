package src

import (
	"strconv"

	"github.com/claudeseo/railgun/src/config"
)

func Init() {
	cfg := config.GetConfig()
	endpoint := cfg.Host + ":" + strconv.Itoa(cfg.Port)

	r := NewRouter()
	r.Run(endpoint)
}
