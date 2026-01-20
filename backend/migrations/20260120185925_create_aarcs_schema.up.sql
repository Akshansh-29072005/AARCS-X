-- Institution Tabel
CREATE TABLE institutions (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);

-- Department Tabel
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT UNIQUE,
    institution_id INT REFERENCES institutions(id)
);

-- Semester Table
CREATE TABLE semesters (
    id SERIAL PRIMARY KEY,
    number INT NOT NULL,
    department_id INT REFERENCES departments(id)
);

-- Subject Table
CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    semester_id INT REFERENCES semesters(id)
);

-- Students Table
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    semester_id INT REFERENCES semesters(id),
    department_id INT REFERENCES departments(id),
    institution_id INT REFERENCES institutions(id),
    role TEXT DEFAULT 'student',
    created_at TIMESTAMP DEFAULT now()
);

-- Teachers
CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    department_id INT REFERENCES departments(id),
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
