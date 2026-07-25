-- name: CreateStudent :execresult
INSERT INTO Student (first_name, last_name, email, password, dob)
VALUES (?, ?, ?, ?, ?);

-- name: GetStudentByEmail :one
SELECT * FROM Student
WHERE email = ? LIMIT 1;

-- name: GetStudentByID :one
SELECT * FROM Student
WHERE student_id = ? LIMIT 1;

-- name: AddStudentAcademic :exec
INSERT INTO Student_Academics (student_id, exam_level, year, roll_no, reg_no, gpa, board)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetStudentAcademics :many
SELECT * FROM Student_Academics
WHERE student_id = ?;

-- name: GetStudentSubjectMarks :many
SELECT subject_name, marks, grade
FROM Student_Subject_Marks
WHERE student_id = ? AND exam_level = ?;

-- name: UpsertStudentProfileField :exec
INSERT INTO Student_Profile_Info (student_id, field_name, field_value)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE field_value = VALUES(field_value);
