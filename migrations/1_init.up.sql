-- =========================
-- TABLE: classes
-- =========================
CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    grade INT NOT NULL CHECK (grade BETWEEN 1 AND 6),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- TABLE: students
-- =========================
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    nisn VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    class_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_students_class
        FOREIGN KEY(class_id)
        REFERENCES classes(id)
        ON DELETE CASCADE
);

-- =========================
-- INDEX (IMPORTANT)
-- =========================

-- Untuk JOIN & filter by class
CREATE INDEX idx_students_class_id ON students(class_id);

-- Untuk search by name (basic search)
CREATE INDEX idx_students_name ON students(name);

-- Untuk filter class berdasarkan grade
CREATE INDEX idx_classes_grade ON classes(grade);

-- =========================
-- FUNCTION: auto update updated_at
-- =========================
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- =========================
-- TRIGGER
-- =========================

-- students
CREATE TRIGGER trg_students_updated_at
BEFORE UPDATE ON students
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- classes
CREATE TRIGGER trg_classes_updated_at
BEFORE UPDATE ON classes
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();