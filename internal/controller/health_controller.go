package controller

import (
	"fmt"
	"net/http"
)

func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is healthy")
}
