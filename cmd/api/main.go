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

// Scope required in the JWT to access GET endpoints
const JWT_SCOPE_READ = "read"

// Scope required in the JWT to access POST endpoints
const JWT_SCOPE_WRITE = "write"

func main() {

	cfg := config.NewConfig()

	repo := repo.NewInMemoryEmployeeRepo()

	// Register the routes and handlers
	h := handlers.NewHandler(repo)

	protectedGetAllEmployees := middleware.JWTMiddleware(http.HandlerFunc(h.GetAllEmployees), cfg, JWT_SCOPE_READ)
	http.Handle("GET /employees", protectedGetAllEmployees)

	protectedGetEmployee := middleware.JWTMiddleware(http.HandlerFunc(h.GetEmployee), cfg, JWT_SCOPE_READ)
	http.Handle("GET /employees/{id}", protectedGetEmployee)

	protectedCreateEmployee := middleware.JWTMiddleware(http.HandlerFunc(h.CreateEmployee), cfg, JWT_SCOPE_WRITE)
	http.Handle("POST /employees", protectedCreateEmployee)

	// Run the server
	var port string = fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, http.DefaultServeMux))

}
