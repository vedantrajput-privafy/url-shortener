package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/joho/godotenv"
	"github.com/vedantrajput-privafy/url-shortener/database"
	"github.com/vedantrajput-privafy/url-shortener/handler"
	"github.com/vedantrajput-privafy/url-shortener/repository"
)

func main() {


	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}


	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	
	defer db.Close()

	rdb, err := database.ConnectRedis()
	if err != nil {
		panic(err)
	}
	defer rdb.Close()

	repo := repository.NewURLRepository(db)
	h := &handler.Handler{
		Repository:  repo,
		RedisClient: rdb,
	}

	http.HandleFunc("/", h.RedirectHandler)
	http.HandleFunc("/shorten", h.ShortenHandler)

	fmt.Println("Server is running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
