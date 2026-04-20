CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    grade VARCHAR(2) NOT NULL
);

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    nisn VARCHAR(10) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    class_id INTEGER NOT NULL,
    
    -- Constraint Relasi
    CONSTRAINT fk_student_class 
        FOREIGN KEY (class_id) 
        REFERENCES classes(id) 
        ON DELETE CASCADE
);

CREATE INDEX idx_students_class_id ON students(class_id);

CREATE INDEX idx_students_nisn ON students(nisn);