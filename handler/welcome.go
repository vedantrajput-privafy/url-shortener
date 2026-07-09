package handler

import (
	"fmt"
	"net/http"
)

// method receiver
//
//	receiver                     parameters
func (h *Handler) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to URL Shortener API")
}
