-- name: InsertUniversity :execresult
INSERT INTO University (u_name, website, location, logo_url)
VALUES (?, ?, ?, ?);

-- name: UpdateUniversity :exec
UPDATE University
SET u_name = ?, website = ?, location = ?, logo_url = ?
WHERE u_id = ?;

-- name: GetUniversityByID :one
SELECT u_id, u_name, website, location, logo_url
FROM University
WHERE u_id = ?;

-- name: DeleteUniversity :exec
DELETE FROM University
WHERE u_id = ?;
