package main

import (
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/server"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
)

func main() {
	util.LoadEnv()
	err := gemini.InitializeGeminiClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	server.Start()
}
