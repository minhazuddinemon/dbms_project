ALTER TABLE Admission_Test
DROP FOREIGN KEY fk_admission_test_program,
DROP INDEX idx_admission_test_program_id,
DROP COLUMN program_id,
DROP COLUMN exam_date;
