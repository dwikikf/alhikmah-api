-- =========================
-- DROP TRIGGER
-- =========================
DROP TRIGGER IF EXISTS trg_set_late_flag ON attendances;
DROP TRIGGER IF EXISTS trg_attendance_updated_at ON attendances;

DROP TRIGGER IF EXISTS trg_students_updated_at ON students;
DROP TRIGGER IF EXISTS trg_classes_updated_at ON classes;

-- =========================
-- DROP FUNCTION
-- =========================
DROP FUNCTION IF EXISTS set_late_flag;
DROP FUNCTION IF EXISTS update_updated_at_column;

-- =========================
-- DROP TABLE
-- =========================
-- urutan: child → parent
DROP TABLE IF EXISTS attendances;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS classes;

-- =========================
-- DROP ENUM
-- =========================
DROP TYPE IF EXISTS attendance_status;
DROP TYPE IF EXISTS attendance_method;