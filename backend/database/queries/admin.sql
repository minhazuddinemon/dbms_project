-- name: GetUniversityApplications :many
SELECT
    a.app_id,
    a.student_id,
    a.program_id,
    p.p_name AS program_name,
    a.status,
    a.sub_date
FROM Application a
JOIN Program p ON a.program_id = p.program_id
WHERE p.u_id = ?
ORDER BY a.sub_date DESC;

-- name: AdminUpdateApplicationStatus :exec
UPDATE Application
SET status = ?
WHERE app_id = ?;
