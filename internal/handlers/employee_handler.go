package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KasperVerhulst/SalaryService/internal/middleware"
	"github.com/KasperVerhulst/SalaryService/internal/models"
	"github.com/KasperVerhulst/SalaryService/internal/repo"
)

type Handler struct {
	employeeRepo repo.EmployeeRepository
	logger       *log.Logger
}

func NewHandler(repo repo.EmployeeRepository) *Handler {
	return &Handler{
		employeeRepo: repo,
		logger:       log.Default(),
	}
}

// Get employees for one company
func (h *Handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	company, _ := r.Context().Value(middleware.UserCtxKey{}).(string)

	h.logger.Println("Getting all employees for company: ", company)
	var employees []models.Employee
	var queryError error

	if company != "" {
		employees, queryError = h.employeeRepo.FindByCompany(company)
		if queryError != nil {
			http.Error(w, "No employees found for your company", http.StatusNotFound)
			return
		}
	} else {
		// no company claim in token
		// get employees for all companies from repo
		employees = h.employeeRepo.FindAll()
	}

	// return all employees as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (h *Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	company, _ := r.Context().Value(middleware.UserCtxKey{}).(string)

	idString := r.PathValue("id")
	h.logger.Println("Getting employee with ID: ", idString)

	// convert string idString to int
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	employee, err := h.employeeRepo.FindByID(id)

	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// if company claim is present in token, check if employee belongs to that company
	if !(company == "") && strings.ToLower(employee.Company) != strings.ToLower(company) {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)

}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	company, _ := r.Context().Value(middleware.UserCtxKey{}).(string)

	h.logger.Println("Adding new employee for company : ", company)

	// serialize request body into employee struct
	var body models.CreateEmployeeRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Creating new employee: %+v\n", body.Name)

	// create new employee
	e := models.Employee{
		ID:       len(h.employeeRepo.FindAll()) + 1, // not very performant but works for in-memory
		Company:  strings.Title(company),
		Name:     body.Name,
		JoinDate: time.Now().Format("2006-01-02"), // set join date to today
		Role:     body.Role,
		Paygrade: body.Paygrade,
	}

	// save employee to repo
	h.employeeRepo.Store(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // set status to 201
	json.NewEncoder(w).Encode(e)
}
