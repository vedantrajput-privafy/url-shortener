package handler

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/vedantrajput-privafy/url-shortener/model"
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

	shortcode := "abc123"

	err = h.Repository.Save(req.OriginalURL, shortcode)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// _, err = h.DB.Exec(query, req.OriginalURL, shortcode)
	// if err != nil {
	// 	fmt.Println("DB Error:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

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
