ALTER TABLE Admission_Test
ADD COLUMN exam_date DATE NULL AFTER exam_center,
ADD COLUMN program_id INT NULL AFTER prereq_test_id,
ADD INDEX idx_admission_test_program_id (program_id),
ADD CONSTRAINT fk_admission_test_program
FOREIGN KEY (program_id) REFERENCES Program (program_id) ON DELETE CASCADE;
