-- name: GetApplicationByID :one
SELECT a.app_id, a.student_id, a.program_id, a.status, p.application_fee, p.p_name AS program_name
FROM Application a
JOIN Program p ON a.program_id = p.program_id
WHERE a.app_id = ? AND a.student_id = ?;

-- name: CreatePayment :execresult
INSERT INTO Payment (app_id, amount, method, tx_id, status)
VALUES (?, ?, ?, ?, 'SUCCESS');

-- name: UpdateApplicationStatus :exec
UPDATE Application
SET status = ?
WHERE app_id = ? AND student_id = ?;
