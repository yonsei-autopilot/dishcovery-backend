package main

import (
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/firebase"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/gemini"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/google_tts"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/logger"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/server"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/util"
)

func main() {
	util.LoadEnv()
	firebase.InitializeFirebaseClient()
	gemini.InitializeGeminiClient()
	google_tts.InitializeGoogleTtsClient()
	logger.InitializeLogger()
	util.InitializeKst()
	server.Start()
}
