package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vedantrajput-privafy/url-shortener/database"
	"github.com/vedantrajput-privafy/url-shortener/utils"
)

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// REMOVE THE LEADING SLASH FROM THE PATH
	code := r.URL.Path[1:]
	id, err := utils.DecodeBase62(code)
	if err != nil {
		http.Error(w, "Invalid shortcode", http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("url:%d", id)

	longURL, err := h.RedisClient.Get(database.Ctx, cacheKey).Result()

	if err == nil {
		// fmt.Println("Redis Hit")
		http.Redirect(w, r, longURL, http.StatusFound)
		return
	}

	if err != redis.Nil {
		http.Error(w, "Redis error", http.StatusInternalServerError)
		return
	}

	// fmt.Println("Redis Miss")

	longURL, err = h.Repository.GetByID(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	err = h.RedisClient.Set(database.Ctx, cacheKey, longURL, 24*time.Hour).Err()

	if err != nil {
		fmt.Println("Redis Set Error", err)
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
