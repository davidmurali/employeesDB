package database

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
type DBInterface interface {
	CreateEmployee(emp *Employee) error
	GetEmployeeByID(id int) (*Employee, error)
	UpdateEmployee(id int) (*Employee, error)
	DeleteEmployee(id int) (*Employee, error)
}
