package main

import (
	"fmt"
	"github.com/vedantrajput-privafy/url-shortener/database"
	"github.com/vedantrajput-privafy/url-shortener/handler"
	"github.com/vedantrajput-privafy/url-shortener/repository"
	"net/http"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewURLRepository(db)
	h := &handler.Handler{
		Repository: repo,
	}

	http.HandleFunc("/", h.WelcomeHandler)
	http.HandleFunc("/shorten", h.ShortenHandler)

	fmt.Println("Server is running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
