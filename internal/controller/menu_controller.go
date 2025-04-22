package controller

import (
	"fmt"
	"net/http"
)

func uploadMenuImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Menu image is uploaded")
}
