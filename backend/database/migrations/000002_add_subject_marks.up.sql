-- Add the education group to the existing table
ALTER TABLE Student_Academics
ADD COLUMN edu_group VARCHAR(50) NOT NULL DEFAULT 'Science';
-- Note: Provide a default temporarily if you already have rows, or truncate the table first.

-- Create the new Subject Marks table
CREATE TABLE Student_Subject_Marks (
    student_id INT NOT NULL,
    exam_level VARCHAR(20) NOT NULL,
    subject_name VARCHAR(50) NOT NULL, -- e.g., 'Physics', 'Mathematics'
    marks NUMERIC(5, 2) NOT NULL,
    grade VARCHAR(5), -- e.g., 'A+', 'A'
    PRIMARY KEY (student_id, exam_level, subject_name),
    FOREIGN KEY (student_id, exam_level) REFERENCES Student_Academics (
        student_id, exam_level
    ) ON DELETE CASCADE
);
