-- query.sql

-- name: CreateTask :one
INSERT INTO tasks (
    title, description
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY created_at DESC;

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET is_completed = $2
WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
