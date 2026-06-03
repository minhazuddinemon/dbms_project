-- name: AddStudentMobile :exec
INSERT INTO
  student_mobiles (student_id, mobile)
VALUES
  ($1, $2);
