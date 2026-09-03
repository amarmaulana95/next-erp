CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    employee_code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(150) NOT NULL,
    department VARCHAR(100),
    position VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO employees (employee_code, name, department, position)
VALUES
    ('EMP001', 'Amar Maulana', 'IT', 'Software Engineer'),
    ('EMP002', 'Budi Santoso', 'Finance', 'Financial Analyst'),
    ('EMP003', 'Citra Lestari', 'HR', 'HR Specialist');
