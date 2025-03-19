package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

type Response struct {
    CurrentTime string `json:"current_time"`
}

func handleRequest(writer http.ResponseWriter, req *http.Request) {
    if req.Method == http.MethodGet {
        data := Response{CurrentTime: time.Now().Format(time.RFC3339)}
        writer.Header().Set("Content-Type", "application/json")
        if encodeErr := json.NewEncoder(writer).Encode(data); encodeErr != nil {
            http.Error(writer, "Error encoding JSON", http.StatusInternalServerError)
            log.Printf("Encoding error: %v", encodeErr)
        }
        return
    }
    http.Error(writer, "Unsupported method", http.StatusMethodNotAllowed)
}

func main() {
    endpoint := "/time"
    port := 8795
    address := fmt.Sprintf(":%d", port)

    http.HandleFunc(endpoint, handleRequest)
    log.Printf("Server running on %s...", address)

    if serverErr := http.ListenAndServe(address, nil); serverErr != nil {
        log.Fatalf("Failed to launch server: %v", serverErr)
    }
}