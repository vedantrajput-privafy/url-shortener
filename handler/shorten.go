package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/vedantrajput-privafy/url-shortener/model"
	"github.com/vedantrajput-privafy/url-shortener/utils"
)

func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, err = url.ParseRequestURI(req.OriginalURL)
	if err != nil {
		http.Error(w, "Please give correct url", http.StatusBadRequest)
		return
	}

	id, err := h.Repository.Save(req.OriginalURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	shortcode := utils.EncodeBase62(id)

	response := model.ShortenResponse{
		ShortURL: "http://localhost:8080/" + shortcode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
