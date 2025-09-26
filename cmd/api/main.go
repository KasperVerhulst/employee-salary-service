package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KasperVerhulst/SalaryService/config"
	"github.com/KasperVerhulst/SalaryService/internal/handlers"
	"github.com/KasperVerhulst/SalaryService/internal/middleware"
	"github.com/KasperVerhulst/SalaryService/internal/repo"
)

func main() {

	cfg := config.NewConfig()

	repo := repo.NewInMemoryEmployeeRepo()

	// Register the routes and handlers
	h := handlers.NewHandler(repo)

	http.HandleFunc("GET /employees", h.GetAllEmployees)
	http.HandleFunc("GET /employees/{id}", h.GetEmployee)
	http.HandleFunc("POST /employees", h.CreateEmployee)

	// Run the server
	var port string = fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, middleware.JWTMiddleware(http.DefaultServeMux, cfg)))

}
