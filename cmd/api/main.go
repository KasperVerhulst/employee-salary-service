package main

import (
	"log"
	"net/http"

	"github.com/KasperVerhulst/SalaryService/internal/handlers"
	"github.com/KasperVerhulst/SalaryService/internal/repo"
)

func main() {

	repo := repo.NewInMemoryEmployeeRepo()

	// Register the routes and handlers
	h := handlers.NewHandler(repo)

	http.HandleFunc("GET /employees", h.GetAllEmployees)
	http.HandleFunc("GET /employees/{id}", h.GetEmployee)
	http.HandleFunc("POST /employees", h.CreateEmployee)

	// Run the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
