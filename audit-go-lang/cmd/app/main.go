package main

import (
	"audit/internal/http"
	"audit/internal/repository"
)

func main() {
	repository.ConnectDatabase()
	http.ListenAndServe()
}
