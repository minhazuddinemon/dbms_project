-- name: ListUniversities :many
SELECT * FROM University;

-- name: ListProgramsByUniversity :many
SELECT * FROM Program
WHERE u_id = ?;

-- name: GetProgramDetails :one
SELECT p.*, u.u_name, u.location
FROM Program p
JOIN University u ON p.u_id = u.u_id
WHERE p.program_id = ? LIMIT 1;

-- name: GetProgramEligibilityRules :many
SELECT * FROM Program_Eligibility_Rules
WHERE program_id = ?;
