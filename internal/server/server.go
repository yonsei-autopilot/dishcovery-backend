package server

import (
	"log"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/controller"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/middleware"
)

func Start() {
	handler := http.NewServeMux()
	controller.RegisterRoutes(handler)

	server := http.Server{
		Addr:    ":8090",
		Handler: middleware.Logging(handler),
	}

	server.ListenAndServe()
	log.Println("Server running")
}
