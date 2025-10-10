package repo

import (
	"errors"
	"strings"

	"github.com/KasperVerhulst/SalaryService/internal/models"
)

type EmployeeRepository interface {
	FindAll() []models.Employee
	FindByID(ID int) (models.Employee, error)
	FindByCompany(company string) ([]models.Employee, error)
	FindByIDAndCompany(ID int, company string) (models.Employee, error)
	Store(e models.Employee) error

	DeleteByID(ID int) error
}

type InMemoryEmployeeRepository struct {
	employees []models.Employee
}

func NewInMemoryEmployeeRepo() *InMemoryEmployeeRepository {
	return &InMemoryEmployeeRepository{
		employees: []models.Employee{
			{ID: 1, Company: "Proximus", Name: "Jan Peeters", JoinDate: "2021-02-15", Role: "Network Engineer", Paygrade: "P3"},
			{ID: 2, Company: "Telenet", Name: "Emma Janssens", JoinDate: "2020-05-20", Role: "Marketing Manager", Paygrade: "M2"},
			{ID: 3, Company: "Telenet", Name: "Wouter Claes", JoinDate: "2019-11-05", Role: "Software Developer", Paygrade: "P2"},
			{ID: 4, Company: "Proximus", Name: "Marie Dubois", JoinDate: "2022-01-10", Role: "Data Analyst", Paygrade: "P1"},
			{ID: 5, Company: "Proximus", Name: "Pieter Mertens", JoinDate: "2018-07-01", Role: "Senior Technician", Paygrade: "T3"},
			{ID: 6, Company: "Telenet", Name: "Julie Lambert", JoinDate: "2023-03-12", Role: "HR Business Partner", Paygrade: "H2"},
			{ID: 7, Company: "Proximus", Name: "Thomas Leroy", JoinDate: "2020-09-08", Role: "Customer Support Agent", Paygrade: "S1"},
			{ID: 8, Company: "Telenet", Name: "Laura Simon", JoinDate: "2021-06-22", Role: "Product Owner", Paygrade: "M1"},
			{ID: 9, Company: "Telenet", Name: "Dries Goossens", JoinDate: "2022-08-30", Role: "Field Technician", Paygrade: "T2"},
			{ID: 10, Company: "Proximus", Name: "Sarah Martin", JoinDate: "2017-04-18", Role: "Sales Director", Paygrade: "M4"},
			{ID: 11, Company: "Telenet", Name: "Michiel Wouters", JoinDate: "2023-02-02", Role: "Junior Developer", Paygrade: "P1"},
			{ID: 12, Company: "Proximus", Name: "Eva Jacobs", JoinDate: "2019-10-14", Role: "Systems Administrator", Paygrade: "P3"},
			{ID: 13, Company: "Proximus", Name: "Robbe Willems", JoinDate: "2021-12-07", Role: "Network Analyst", Paygrade: "P2"},
			{ID: 14, Company: "Telenet", Name: "Lisa Dupont", JoinDate: "2022-07-19", Role: "Graphic Designer", Paygrade: "C2"},
			{ID: 15, Company: "Proximus", Name: "Louis Maes", JoinDate: "2020-03-03", Role: "Finance Manager", Paygrade: "M3"},
		},
	}
}

func (r *InMemoryEmployeeRepository) FindAll() []models.Employee {
	return r.employees
}

func (r *InMemoryEmployeeRepository) FindByID(ID int) (models.Employee, error) {
	for _, e := range r.employees {
		if e.ID == ID {
			return e, nil
		}
	}
	return models.Employee{}, errors.New("employee not found")
}

func (r *InMemoryEmployeeRepository) FindByCompany(company string) ([]models.Employee, error) {
	var result []models.Employee
	for _, e := range r.employees {
		if strings.ToLower(e.Company) == strings.ToLower(company) {
			result = append(result, e)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("no employees found for company")
	}
	return result, nil
}

func (r *InMemoryEmployeeRepository) FindByIDAndCompany(ID int, company string) (models.Employee, error) {
	for _, e := range r.employees {
		if e.ID == ID && e.Company == company {
			return e, nil
		}
	}
	return models.Employee{}, errors.New("employee not found for company")
}

func (r *InMemoryEmployeeRepository) Store(e models.Employee) error {
	r.employees = append(r.employees, e)
	return nil
}

func (r *InMemoryEmployeeRepository) DeleteByID(ID int) error {
	for i, e := range r.employees {
		if e.ID == ID {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}
	return errors.New("employee not found")
}

func (r *InMemoryEmployeeRepository) Update(e models.Employee) error {
	return errors.New("employee not found")
}
