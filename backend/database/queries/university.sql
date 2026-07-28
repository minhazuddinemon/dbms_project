-- name: InsertUniversity :execresult
INSERT INTO University (u_name, website, location, logo_url, university_description, university_history)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateUniversity :exec
UPDATE University
SET u_name = ?, website = ?, location = ?, logo_url = ?, university_description = ?, university_history = ?
WHERE u_id = ?;

-- name: GetUniversityByID :one
SELECT u_id, u_name, website, location, logo_url, university_description, university_history
FROM University
WHERE u_id = ?;

-- name: DeleteUniversity :exec
DELETE FROM University
WHERE u_id = ?;

-- name: InsertUniversityDepartment :execresult
INSERT INTO University_Department (u_id, dept_name, dept_description, total_seats)
VALUES (?, ?, ?, ?);

-- name: DeleteUniversityDepartments :exec
DELETE FROM University_Department
WHERE u_id = ?;

-- name: GetUniversityDepartments :many
SELECT dept_id, u_id, dept_name, dept_description, total_seats
FROM University_Department
WHERE u_id = ?;

-- name: InsertUniversityAlbumPicture :execresult
INSERT INTO University_Album (u_id, picture_title, picture_url)
VALUES (?, ?, ?);

-- name: DeleteUniversityAlbum :exec
DELETE FROM University_Album
WHERE u_id = ?;

-- name: GetUniversityAlbum :many
SELECT album_id, u_id, picture_title, picture_url
FROM University_Album
WHERE u_id = ?;

