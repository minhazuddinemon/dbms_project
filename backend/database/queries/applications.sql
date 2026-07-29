-- name: CreateApplication :execresult
INSERT INTO Application (program_id, student_id, status)
VALUES (?, ?, 'Pending');

-- name: GetStudentApplications :many
SELECT
    a.app_id,
    a.program_id,
    a.sub_date,
    a.status,
    p.p_name AS program_name,
    u.u_name AS university_name
FROM Application a
JOIN Program p ON a.program_id = p.program_id
JOIN University u ON p.u_id = u.u_id
WHERE a.student_id = ?
ORDER BY a.sub_date DESC;

-- name: RecordPayment :execresult
INSERT INTO Payment (tx_id, amount, status, method, app_id)
VALUES (?, ?, ?, ?, ?);

-- name: GetProgramRequiredFields :many
SELECT field_name
FROM Program_Required_Fields
WHERE program_id = ?;

-- name: GetStudentProfileFields :many
SELECT field_name
FROM Student_Profile_Info
WHERE student_id = ?;

-- name: CheckExistingApplication :one
SELECT app_id
FROM Application
WHERE student_id = ? AND program_id = ? LIMIT 1;
