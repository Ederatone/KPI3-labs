package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC3339)
		response := TimeResponse{
			Time: currentTime,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Сервер запущений на http://localhost:8795")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		fmt.Println("Помилка при запуску сервера:", err)
	}
}
