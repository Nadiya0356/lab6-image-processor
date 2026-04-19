package main

import (
	"lab6-image-processor/internal/api"
	"lab6-image-processor/internal/processor"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on :8080")

	storage := &processor.MemoryStorage{}

	go processor.RunWorkerPool(5, storage)

	http.HandleFunc("/health", api.HealthHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
