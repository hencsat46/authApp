package main

import (
	db "authServ/internal/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func parseURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		w.Write([]byte("Кнопка нажата"))
		break
	case "POST":

		decodedJson := make(map[string]string)
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&decodedJson)
		fmt.Println(decodedJson)

		switch decodedJson["type"] {
		case "Sign up":
			db.AddEncodeData(decodedJson)
			break
		case "Sign in":
			err := db.Authorization(decodedJson)
			fmt.Println(err)
			break
		}

		if decodedJson["type"] == "Sign up" {

		}

		//
		break
	}

}

func main() {

	r := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	r.HandleFunc("/", parseURL)
	handler := c.Handler(r)
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":3000", handler))

}
