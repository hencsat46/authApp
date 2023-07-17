package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func parseURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write([]byte("я сосал меня ебали"))
		break
	case "POST":
		w.Write([]byte("я сосал меня ебали"))
		break
	}

}

func main() {

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	r.HandleFunc("/", parseURL)
	handler := c.Handler(r)
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":5050", handler))

}
