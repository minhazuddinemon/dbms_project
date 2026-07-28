-- name: RecordTestResult :exec
INSERT INTO Gives (student_id, test_id, marks, merit_position)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE marks = VALUES(marks), merit_position = VALUES(merit_position);

-- name: GetStudentTestResults :many
SELECT g.marks, g.merit_position, t.exam_unit, u.u_name, p.p_name AS program_name
FROM Gives g
JOIN Admission_Test t ON g.test_id = t.test_id
LEFT JOIN Program p ON t.program_id = p.program_id
JOIN Conducts c ON t.test_id = c.test_id
JOIN University u ON c.u_id = u.u_id
WHERE g.student_id = ?;

-- name: GetAdmissionTestByProgramID :one
SELECT test_id, exam_unit, exam_center, exam_date, prereq_test_id, program_id
FROM Admission_Test
WHERE program_id = ? LIMIT 1;

-- name: GetAdmissionTestByID :one
SELECT test_id, exam_unit, exam_center, exam_date, prereq_test_id, program_id
FROM Admission_Test
WHERE test_id = ? LIMIT 1;
