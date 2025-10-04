package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})
	println("Server starting on http://0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
