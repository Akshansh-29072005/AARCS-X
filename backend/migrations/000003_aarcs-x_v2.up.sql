-- Role enums scoped to academic context.
CREATE TYPE faculty_role AS ENUM ('faculty', 'hod');
CREATE TYPE student_role AS ENUM ('student', 'cr');
CREATE TYPE assessment_type AS ENUM ('CT1', 'CT2');

-- User identity and authentication.
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Institution root entity.
CREATE TABLE institutions (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    official_email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT NOT NULL,
    district TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Exactly one owner record per institution.
CREATE TABLE institution_owners (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    institution_id BIGINT NOT NULL UNIQUE REFERENCES institutions(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Institution-level admins appointed by owner.
CREATE TABLE institution_admins (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    institution_id BIGINT NOT NULL REFERENCES institutions(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, institution_id)
);

-- Departments belong to one institution.
CREATE TABLE departments (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(50) NOT NULL,
    institution_id BIGINT NOT NULL REFERENCES institutions(id) ON DELETE CASCADE,
    head_of_department_faculty_id BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (institution_id, code)
);

-- Faculties are institution-scoped and department-scoped memberships.
CREATE TABLE faculties (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    institution_id BIGINT NOT NULL REFERENCES institutions(id) ON DELETE CASCADE,
    department_id BIGINT NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    role faculty_role NOT NULL DEFAULT 'faculty',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, institution_id),
    UNIQUE (user_id, department_id)
);

ALTER TABLE departments
    ADD CONSTRAINT fk_departments_head_of_department
    FOREIGN KEY (head_of_department_faculty_id)
    REFERENCES faculties(id)
    ON DELETE SET NULL;

-- Semesters are department-scoped.
CREATE TABLE semesters (
    id BIGSERIAL PRIMARY KEY,
    number INT NOT NULL CHECK (number >= 1 AND number <= 12),
    department_id BIGINT NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (department_id, number),
    UNIQUE (id, department_id)
);

-- Subjects are semester-scoped with code uniqueness per semester.
CREATE TABLE subjects (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(50) NOT NULL,
    semester_id BIGINT NOT NULL REFERENCES semesters(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (semester_id, code)
);

-- Teachers to subjects assignment map.
CREATE TABLE faculty_subjects (
    id BIGSERIAL PRIMARY KEY,
    faculty_id BIGINT NOT NULL REFERENCES faculties(id) ON DELETE CASCADE,
    subject_id BIGINT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (faculty_id, subject_id)
);

-- Students are institution-scoped memberships tied to one department and semester.
CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    institution_id BIGINT NOT NULL REFERENCES institutions(id) ON DELETE CASCADE,
    department_id BIGINT NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    semester_id BIGINT NOT NULL,
    role student_role NOT NULL DEFAULT 'student',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, institution_id),
    FOREIGN KEY (semester_id, department_id)
        REFERENCES semesters(id, department_id)
        ON DELETE RESTRICT
);

-- One attendance row per student-subject-day, marked by faculty.
CREATE TABLE attendance (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    subject_id BIGINT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status BOOLEAN NOT NULL,
    marked_by BIGINT NOT NULL REFERENCES faculties(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (student_id, subject_id, date)
);

-- Assessment records for students by subject and test type.
CREATE TABLE assessments (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    subject_id BIGINT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    test_type assessment_type NOT NULL,
    marks INT NOT NULL CHECK (marks >= 0),
    max_marks INT NOT NULL CHECK (max_marks > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (marks <= max_marks),
    UNIQUE (student_id, subject_id, test_type)
);

-- Performance indexes for joins/filtering.
CREATE INDEX idx_departments_institution_id ON departments(institution_id);
CREATE INDEX idx_faculties_institution_id ON faculties(institution_id);
CREATE INDEX idx_faculties_department_id ON faculties(department_id);
CREATE INDEX idx_semesters_department_id ON semesters(department_id);
CREATE INDEX idx_subjects_semester_id ON subjects(semester_id);
CREATE INDEX idx_students_institution_id ON students(institution_id);
CREATE INDEX idx_students_department_id ON students(department_id);
CREATE INDEX idx_students_semester_id ON students(semester_id);
CREATE INDEX idx_attendance_student_id ON attendance(student_id);
CREATE INDEX idx_attendance_subject_id ON attendance(subject_id);
CREATE INDEX idx_assessments_student_id ON assessments(student_id);
CREATE INDEX idx_assessments_subject_id ON assessments(subject_id);

