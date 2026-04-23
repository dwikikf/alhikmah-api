-- =========================
-- DROP TRIGGER
-- =========================
DROP TRIGGER IF EXISTS trg_students_updated_at ON students;
DROP TRIGGER IF EXISTS trg_classes_updated_at ON classes;

-- =========================
-- DROP FUNCTION
-- =========================
DROP FUNCTION IF EXISTS update_updated_at_column;

-- =========================
-- DROP TABLE
-- =========================
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS classes;