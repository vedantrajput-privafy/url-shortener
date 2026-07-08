package main

import(
	"fmt"
	"net/http"
	"github.com/vedantrajput-privafy/url-shortener/handler"
	"github.com/vedantrajput-privafy/url-shortener/database"
)

func main(){

	db,err := database.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	h:= &handler.Handler{
		DB : db,
	}

	http.HandleFunc("/", h.WelcomeHandler)
	http.HandleFunc("/shorten", h.ShortenHandler)


	fmt.Println("Server is running on http://localhost:8080")

	err = http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Println("Error starting the server:",err)
	}
}