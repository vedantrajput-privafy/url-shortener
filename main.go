package main

import(
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Welcome to URL Shortener API")
}

func main(){
	http.HandleFunc("/",welcome)
	
	fmt.Println("Server is running on http://localhost:8080")

	err:= http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Println("Error starting the server:",err)
	}
}