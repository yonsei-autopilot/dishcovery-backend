package main

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/firebase"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/server"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
)

func main() {
	util.LoadEnv()
	firebase.InitializeFirebaseClient()
	gemini.InitializeGeminiClient()
	server.Start()
}
