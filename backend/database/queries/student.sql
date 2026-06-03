-- name: CreateStudent :one
INSERT INTO
  students (
    first_name,
    last_name,
    email,
    password,
    date_of_birth,
    ssc_roll,
    ssc_reg,
    ssc_gpa,
    ssc_board,
    ssc_year,
    hsc_roll,
    hsc_reg,
    hsc_gpa,
    hsc_board,
    hsc_year
  )
VALUES
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
  )
RETURNING
  *;

-- name: GetStudent :one
SELECT
  *
FROM
  students
WHERE
  student_id = $1
LIMIT
  1;
