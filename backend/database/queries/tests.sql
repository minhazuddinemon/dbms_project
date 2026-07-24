-- name: RecordTestResult :exec
INSERT INTO Gives (student_id, test_id, marks, merit_position)
VALUES (?, ?, ?, ?);

-- name: GetStudentTestResults :many
SELECT g.marks, g.merit_position, t.exam_unit, u.u_name
FROM Gives g
JOIN Admission_Test t ON g.test_id = t.test_id
JOIN Conducts c ON t.test_id = c.test_id
JOIN University u ON c.u_id = u.u_id
WHERE g.student_id = ?;
