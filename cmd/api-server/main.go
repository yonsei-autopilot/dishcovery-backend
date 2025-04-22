package main

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/config"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/server"
)

func main() {
	config.LoadEnv()
	server.Start()
}
