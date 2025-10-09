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

const (
	JWT_SCOPE_READ   string = "read"   // Scope required in the JWT to access GET endpoints
	JWT_SCOPE_WRITE  string = "write"  // Scope required in the JWT to access POST/PUT endpoints
	JWT_SCOPE_DELETE string = "delete" // Scope required in the JWT to access DELETE endpoints
)

const API_RESOURCE_ROOT_PATH string = "/employees"

func main() {

	cfg := config.NewConfig()

	repo := repo.NewInMemoryEmployeeRepo()

	// Register the routes and handlers
	h := handlers.NewHandler(repo)

	// GET /employees
	var getAllEmployeesPath string = fmt.Sprintf("GET %s", API_RESOURCE_ROOT_PATH)
	protectedGetAllEmployees := middleware.JWTMiddleware(http.HandlerFunc(h.GetAllEmployees), cfg, JWT_SCOPE_READ)
	http.Handle(getAllEmployeesPath, protectedGetAllEmployees)

	// GET /employees/{id}
	var getEmployeePath string = fmt.Sprintf("GET %s/{id}", API_RESOURCE_ROOT_PATH)
	protectedGetEmployee := middleware.JWTMiddleware(http.HandlerFunc(h.GetEmployee), cfg, JWT_SCOPE_READ)
	http.Handle(getEmployeePath, protectedGetEmployee)

	// POST /employees
	var postEmployeePath string = fmt.Sprintf("POST %s", API_RESOURCE_ROOT_PATH)
	protectedCreateEmployee := middleware.JWTMiddleware(http.HandlerFunc(h.CreateEmployee), cfg, JWT_SCOPE_WRITE)
	http.Handle(postEmployeePath, protectedCreateEmployee)

	// PUT /employees/{id}
	var putEmployeePath string = fmt.Sprintf("PUT %s/{id}", API_RESOURCE_ROOT_PATH)
	protectedUpdateEmployee := middleware.JWTMiddleware(http.HandlerFunc(h.UpdateEmployee), cfg, JWT_SCOPE_WRITE)
	http.Handle(putEmployeePath, protectedUpdateEmployee)

	// PUT /employees
	var putBatchEmployeePath string = fmt.Sprintf("PUT %s", API_RESOURCE_ROOT_PATH)
	http.HandleFunc(putBatchEmployeePath, h.UpdateBatchEmployees)

	// DELETE /employees/{id}
	var deleteEmployeePath string = fmt.Sprintf("DELETE %s/{id}", API_RESOURCE_ROOT_PATH)
	protectedDeleteEmployee := middleware.JWTMiddleware(http.HandlerFunc(h.DeleteEmployee), cfg, JWT_SCOPE_DELETE)
	http.Handle(deleteEmployeePath, protectedDeleteEmployee)

	// DELETE /employees
	var deleteAllEmployeesPath string = fmt.Sprintf("DELETE %s", API_RESOURCE_ROOT_PATH)
	http.HandleFunc(deleteAllEmployeesPath, h.DeleteAllEmployees)

	// Run the server
	var port string = fmt.Sprintf(":%d", cfg.Port)
	log.Fatal(http.ListenAndServe(port, http.DefaultServeMux))

}
