-- name: GetTodo :one
SELECT
  *
FROM
  todo
WHERE
  id = ?
LIMIT
  1;

-- name: GetTodos :many
SELECT
  *
FROM
  todo
ORDER BY
  created_at;

-- name: CreateTodo :one
INSERT INTO
  todo (content)
VALUES
  (?) RETURNING *;

-- name: ToggleTodo :exec
UPDATE
  todo
SET
  is_finished = ?
WHERE
  id = ?;

-- name: DeleteTodo :exec
DELETE FROM
  todo
WHERE
  id = ?