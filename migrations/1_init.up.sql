-- =========================
-- ENUM TYPES
-- =========================
CREATE TYPE attendance_status AS ENUM (
    'hadir',
    'izin',
    'sakit',
    'alpha'
);

CREATE TYPE attendance_method AS ENUM (
    'qr',
    'manual',
    'admin'
);

-- =========================
-- TABLE: classes
-- =========================
CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    grade INT NOT NULL CHECK (grade BETWEEN 1 AND 6),
    start_time TIME DEFAULT '07:30:00',
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
-- TABLE: attendances
-- =========================
CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,

    attendance_date DATE NOT NULL DEFAULT CURRENT_DATE,

    check_in TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    check_out TIMESTAMP NULL,

    status attendance_status NOT NULL DEFAULT 'hadir',
    method attendance_method NOT NULL DEFAULT 'qr',
    note TEXT,

    is_late BOOLEAN,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_attendance_student
        FOREIGN KEY (student_id)
        REFERENCES students(id)
        ON DELETE RESTRICT
        ON UPDATE RESTRICT
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

-- Untuk filter kehadiran berdasarkan status, tanggal, dan keterlambatan
CREATE INDEX idx_attendance_student_id ON attendances(student_id);
CREATE INDEX idx_attendance_date ON attendances(attendance_date);
CREATE INDEX idx_attendance_status ON attendances(status);
CREATE INDEX idx_attendance_late ON attendances(is_late);

 -- Untuk memastikan satu siswa hanya bisa memiliki satu catatan kehadiran per hari
CREATE UNIQUE INDEX unique_student_per_day
ON attendances(student_id, attendance_date);

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

-- function untuk set is_late berdasarkan check_in dan start_time kelas
CREATE OR REPLACE FUNCTION set_late_flag()
RETURNS TRIGGER AS $$
DECLARE
    class_start TIME;
BEGIN
    SELECT c.start_time INTO class_start
    FROM students s
    JOIN classes c ON s.class_id = c.id
    WHERE s.id = NEW.student_id;

    IF NEW.is_late IS NULL THEN
        IF NEW.check_in::time > class_start THEN
            NEW.is_late = TRUE;
        ELSE
            NEW.is_late = FALSE;
        END IF;
    END IF;

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

-- attendances
CREATE TRIGGER trg_attendance_updated_at
BEFORE UPDATE ON attendances
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- set is_late flag
CREATE TRIGGER trg_set_late_flag
BEFORE INSERT OR UPDATE OF check_in ON attendances
FOR EACH ROW
EXECUTE FUNCTION set_late_flag();