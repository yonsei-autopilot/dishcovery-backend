package server

import (
	"log"
	"net/http"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/middleware"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/controller"
)

func Start() {
	handler := http.NewServeMux()
	controller.RegisterRoutes(handler)

	server := http.Server{
		Addr:    ":8090",
		Handler: middleware.Logging(handler),
	}

	log.Printf("Server listening on %s\n", server.Addr)
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Println("Server exited with error")
		log.Panicln(err)
	}
}
