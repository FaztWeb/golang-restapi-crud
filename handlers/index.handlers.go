package handlers

import (
	"fmt"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}
