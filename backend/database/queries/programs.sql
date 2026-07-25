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

-- name: ListPrograms :many
SELECT
    p.program_id,
    p.p_name,
    p.p_unit,
    p.total_seats,
    p.prev_cutmarks,
    p.deadline,
    u.u_id,
    u.u_name AS university_name,
    u.location AS university_location
FROM Program p
JOIN University u ON p.u_id = u.u_id
WHERE
    (p.p_name LIKE CONCAT('%', ?, '%') OR u.u_name LIKE CONCAT('%', ?, '%'))
    AND (p.p_unit = ? OR ? = '')
ORDER BY p.program_id DESC;

-- name: GetProgramByID :one
SELECT
    p.program_id,
    p.p_name,
    p.p_unit,
    p.total_seats,
    p.prev_cutmarks,
    p.deadline,
    u.u_id,
    u.u_name AS university_name,
    u.website,
    u.location
FROM Program p
JOIN University u ON p.u_id = u.u_id
WHERE p.program_id = ? LIMIT 1;

-- name: GetAllProgramsWithRules :many
SELECT
    p.program_id,
    p.p_name,
    u.u_name AS university_name,
    r.rule_type,
    r.rule_value
FROM Program p
JOIN University u ON p.u_id = u.u_id
LEFT JOIN Program_Eligibility_Rules r ON p.program_id = r.program_id;
