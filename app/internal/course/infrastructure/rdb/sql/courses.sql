-- name: CreateCourses :execlastid
INSERT INTO courses (id, name, level) VALUES (?, ?, ?);

-- name: ListCourses :many
SELECT id, name, level FROM courses;

-- name: SearchCourses :many
SELECT id, name, level FROM courses WHERE level = ?;
