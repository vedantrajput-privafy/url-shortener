package handler

import (
	"encoding/json"
	"fmt"
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

	// shortcode, err := utils.GenerateShortCode(6)
	// if err != nil {
	// 	http.Error(w, "Failed to generate code", http.StatusInternalServerError)
	// 	return
	// }

	// checking valid url

	// if req.OriginalURL == ""{
	// 	http.Error(w,"Please input valid URL",http.StatusBadRequest)
	// 	return
	// }

	_, err = url.ParseRequestURI(req.OriginalURL)
	if err != nil {
		http.Error(w, "Please give correct url", http.StatusBadRequest)
		return
	}

	var shortcode string

	for {
		shortcode, err = utils.GenerateShortCode(6)
		if err != nil {
			http.Error(w, "Failed to generate code", http.StatusInternalServerError)
			return
		}
		// exists, err := h.Repository.Exists(shortcode)
		// if err != nil {
		// 	http.Error(w, "Database Error", http.StatusInternalServerError)
		// 	return
		// }

		exists, err := h.Repository.Exists(shortcode)
		if err != nil {
			fmt.Println("Exists Error:", err)
			http.Error(w, "Database Error", http.StatusInternalServerError)
			return
		}

		if !exists {
			break
		}
	}

	// fmt.Println(shortcode)

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
