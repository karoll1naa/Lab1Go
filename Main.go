package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC3309)
		response := TimeResponse{Time: currentTime}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Heder().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	http.ListenAndServe(":8711", nil)
}
