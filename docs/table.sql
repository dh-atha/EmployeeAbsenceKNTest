-- Departments Table
CREATE TABLE departments (
    department_id SERIAL PRIMARY KEY,
    department_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP
);

-- Positions Table
CREATE TABLE positions (
    position_id SERIAL PRIMARY KEY,
    department_id INTEGER REFERENCES departments(department_id),
    position_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP
);

-- Locations Table
CREATE TABLE locations (
    location_id SERIAL PRIMARY KEY,
    location_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP
);

-- Employees Table
CREATE SEQUENCE employee_code_seq START 1;

CREATE TABLE employees (
    employee_id SERIAL PRIMARY KEY,
    employee_code VARCHAR(10) UNIQUE NOT NULL DEFAULT (TO_CHAR(CURRENT_DATE, 'DDMM') || LPAD(nextval('employee_code_seq')::TEXT, 3, '0')),
    employee_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    department_id INTEGER REFERENCES departments(department_id),
    position_id INTEGER REFERENCES positions(position_id),
    superior INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP
);

-- Attendance Table
CREATE TABLE attendances (
    attendance_id SERIAL PRIMARY KEY,
    employee_id INTEGER REFERENCES employees(employee_id),
    location_id INTEGER REFERENCES locations(location_id),
    absent_in TIMESTAMP,
    absent_out TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP
);

-- dummy data
-- Insert into department table
INSERT INTO departments (
    department_id,
    department_name,
    created_at,
    created_by,
    updated_at,
    updated_by,
    deleted_at
) VALUES (
    DEFAULT,
    'Human Resources',
    '2024-06-20 10:00:00',
    'admin',
    '2024-06-20 10:00:00',
    'admin',
    NULL
);

-- Insert into position table
INSERT INTO positions (
    position_id,
    department_id,
    position_name,
    created_at,
    created_by,
    updated_at,
    updated_by,
    deleted_at
) VALUES (
    DEFAULT,
    1,  -- Assuming the department_id from the previous insert
    'Manager',
    '2024-06-20 10:00:00',
    'admin',
    '2024-06-20 10:00:00',
    'admin',
    NULL
);

-- Insert into location table
INSERT INTO locations (
    location_id,
    location_name,
    created_at,
    created_by,
    updated_at,
    updated_by,
    deleted_at
) VALUES (
    DEFAULT,
    'Head Office',
    '2024-06-20 10:00:00',
    'admin',
    '2024-06-20 10:00:00',
    'admin',
    NULL
);

-- Insert into employee table
INSERT INTO employees (
    employee_id,
    employee_code,
    employee_name,
    password,
    department_id,
    position_id,
    superior,
    created_at,
    created_by,
    updated_at,
    updated_by,
    deleted_at
) VALUES (
    DEFAULT,
    DEFAULT,
    'John Doe',
    'encrypted_password',
    1,  -- Assuming the department_id from the previous insert
    1,  -- Assuming the position_id from the previous insert
    NULL,
    '2024-06-20 10:00:00',
    'admin',
    '2024-06-20 10:00:00',
    'admin',
    NULL
);