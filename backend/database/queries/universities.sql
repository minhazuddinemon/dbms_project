-- name: CreateUniversity :one
INSERT INTO
  universities (
    university_name,
    website,
    location_address,
    transport_routes,
    est_travel_time
  )
VALUES
  ($1, $2, $3, $4, $5)
RETURNING
  *;

-- name: GetUniversity :one
SELECT
  *
FROM
  universities
WHERE
  university_id = $1
LIMIT
  1;

-- name: ListUniversities :many
SELECT
  *
FROM
  universities
ORDER BY
  university_name;
