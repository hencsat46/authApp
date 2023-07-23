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

		//fmt.Printf("%T", r.Body)
		decodedJson := make(map[string]string)
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&decodedJson)
		fmt.Println(decodedJson)

		db.AddEncodeData(decodedJson)
		db.ConnectDB()
		//fmt.Println(string(result))
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
