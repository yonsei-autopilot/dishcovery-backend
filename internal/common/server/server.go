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

	wrappedHandler := middleware.Logging(middleware.Authentication(handler))

	server := http.Server{
		Addr:    ":8090",
		Handler: wrappedHandler,
	}

	log.Printf("Server listening on %s\n", server.Addr)
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Println("Server exited with error")
		log.Panicln(err)
	}
}
