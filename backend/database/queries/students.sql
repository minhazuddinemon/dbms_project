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
INSERT INTO Student_Academics (student_id, exam_level, year, roll_no, reg_no, gpa, board, edu_group)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE year = VALUES(year), roll_no = VALUES(roll_no), reg_no = VALUES(reg_no), gpa = VALUES(gpa), board = VALUES(board), edu_group = VALUES(edu_group);

-- name: GetStudentAcademics :many
SELECT * FROM Student_Academics
WHERE student_id = ?;

-- name: UpsertStudentSubjectMark :exec
INSERT INTO Student_Subject_Marks (student_id, exam_level, subject_name, marks, grade)
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE marks = VALUES(marks), grade = VALUES(grade);

-- name: GetStudentSubjectMarks :many
SELECT subject_name, marks, grade
FROM Student_Subject_Marks
WHERE student_id = ? AND exam_level = ?;

-- name: UpsertStudentProfileField :exec
INSERT INTO Student_Profile_Info (student_id, field_name, field_value)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE field_value = VALUES(field_value);

-- name: GetStudentMobiles :many
SELECT student_id, mobile_no, owner_type
FROM Student_Mobile
WHERE student_id = ?;

-- name: InsertStudentMobile :execresult
INSERT INTO Student_Mobile (student_id, mobile_no, owner_type)
VALUES (?, ?, ?);

-- name: UpdateStudentMobile :execresult
UPDATE Student_Mobile
SET mobile_no = ?, owner_type = ?
WHERE student_id = ? AND mobile_no = ?;

-- name: DeleteStudentMobile :execresult
DELETE FROM Student_Mobile
WHERE student_id = ? AND mobile_no = ?;

-- name: CreateNotification :execresult
INSERT INTO Notification (student_id, message)
VALUES (?, ?);

-- name: GetStudentNotifications :many
SELECT notif_id, student_id, message, created_at
FROM Notification
WHERE student_id = ?
ORDER BY created_at DESC;
