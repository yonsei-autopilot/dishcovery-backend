package main

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/server"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
)

func main() {
	util.LoadEnv()
	server.Start()
}
