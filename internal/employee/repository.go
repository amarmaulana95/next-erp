package employee

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]Employee, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			employee_code,
			name,
			department,
			position
		FROM employees
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("query employees: %w", err)
	}
	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var employee Employee

		if err := rows.Scan(
			&employee.ID,
			&employee.EmployeeCode,
			&employee.Name,
			&employee.Department,
			&employee.Position,
		); err != nil {
			return nil, fmt.Errorf("scan employee: %w", err)
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate employees: %w", err)
	}

	return employees, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*Employee, error) {
	var employee Employee

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			employee_code,
			name,
			department,
			position
		FROM employees
		WHERE id = $1
	`, id).Scan(
		&employee.ID,
		&employee.EmployeeCode,
		&employee.Name,
		&employee.Department,
		&employee.Position,
	)

	if err != nil {
		return nil, fmt.Errorf("get employee by id: %w", err)
	}

	return &employee, nil
}
