-- Institution Table
CREATE TABLE institutions (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(50) UNIQUE,
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT now()
);

-- Department Table
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(50) UNIQUE,
    head_of_department TEXT NOT NULL,
    institution_id INT REFERENCES institutions(id),
    created_at TIMESTAMP DEFAULT now()
);

-- Semester Table
CREATE TABLE semesters (
    id SERIAL PRIMARY KEY,
    number INT NOT NULL,
    department_id INT REFERENCES departments(id)
    created_at TIMESTAMP DEFAULT now()
);

-- Subject Table
CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    semester_id INT REFERENCES semesters(id)
    created_at TIMESTAMP DEFAULT now()
);

-- Students Table
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL ,
    password VARCHAR(255) NOT NULL,
    semester_id INT REFERENCES semesters(id),
    department_id INT REFERENCES departments(id),
    institution_id INT REFERENCES institutions(id),
    role TEXT DEFAULT 'student',
    created_at TIMESTAMP DEFAULT now()
);

-- Teachers
CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    department_id INT REFERENCES departments(id),
    designation VARCHAR(50) NOT NULL,
    role TEXT DEFAULT 'teacher',
    created_at TIMESTAMP DEFAULT now()
);

-- Teacher-Subject Mapping
CREATE TABLE teacher_subjects (
    id SERIAL PRIMARY KEY,
    teacher_id INT REFERENCES teachers(id),
    subject_id INT REFERENCES subjects(id)
);

-- Attendance Table
CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES students(id),
    subject_id INT REFERENCES subjects(id),
    date DATE NOT NULL,
    status BOOLEAN NOT NULL,
    marked_by INT REFERENCES teachers(id),
    created_at TIMESTAMP DEFAULT now()
);

-- Assessments Table
CREATE TABLE assessments (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES students(id),
    subject_id INT REFERENCES subjects(id),
    test_type TEXT CHECK (test_type IN ('CT1','CT2')),
    marks INT,
    max_marks INT,
    created_at TIMESTAMP DEFAULT now()
);
