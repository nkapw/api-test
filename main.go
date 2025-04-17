package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Menyeting handler untuk URL path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is a simple API server!"))
	})

	// Menyeting handler untuk endpoint /api/apps
	http.HandleFunc("/api/apps", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[{"id": 1, "name": "App One"}, {"id": 2, "name": "App Two"}]`))
	})

	// Jalankan server di port 8080
	fmt.Println("Server is running at :9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
