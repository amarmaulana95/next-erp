package employee

type Employee struct {
	ID           int64  `json:"id"`
	EmployeeCode string `json:"employee_code"`
	Name         string `json:"name"`
	Department   string `json:"department"`
	Position     string `json:"position"`
}
