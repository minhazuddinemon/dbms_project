-- name: GetApplicationByID :one
SELECT app_id, student_id, program_id, status
FROM Application
WHERE app_id = ? AND student_id = ?;

-- name: CreatePayment :execresult
INSERT INTO Payment (app_id, amount, method, tx_id, status)
VALUES (?, ?, ?, ?, 'SUCCESS');

-- name: UpdateApplicationStatus :exec
UPDATE Application
SET status = ?
WHERE app_id = ? AND student_id = ?;
