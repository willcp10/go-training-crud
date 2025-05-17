package main

import (
	"go-training-crud/internal/adapters"
	"log"
	"net/http"
)

func main() {
	adapters.BuildApp()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
