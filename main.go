package main

import (
	"log"
	"net/http"
	"testApp/config"
	"testApp/controllers/userControllers"
)

func main() {
	config.Koneksi()
	http.HandleFunc("/api/user", userControllers.GetDataUser)

	log.Println("Server berjalan di port 8080")
	http.ListenAndServe(":8080", nil)
}
