-- CreateUser :one
INSERT INTO users (
    id,
    email,
    name,
    password,
    status,
    created_at,
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING
    id,
    email,
    name,
    status,
    created_at;


-- GetUserByID :one
SELECT id,
       email,
       name,
       status,
       created_at
FROM users WHERE id = $1;

-- GetUserByEmail :one
SELECT id,
       email,
       name,
       status,
       created_at
FROM users WHERE email = $1;

-- UpdateUser :one
UPDATE users
    name = COALESCE($2, name),
    email = COALESCE($3, email)
WHERE id = $1
RETURNING
    id,
    email,
    name,
    status,
    created_at;

-- DeleteUser :exec
UPDATE users
    status = 0
WHERE id = $1
RETURNING
    id,
    email,
    name,
    status,
    created_at;

-- CheckEmailExist :one
SELECT EXISTS(
    SELECT 1 FROM users WHERE email = $1
) as exists;

-- UpdateUserStatus :exec
UPDATE users
    status = $2
WHERE id = $1
RETURNING
    id,
    email,
    name,
    status,
    created_at;

-- SearchUsers :many
SELECT
    id,
    email,
    name,
    status,
    created_at
FROM users
WHERE email ILIKE '%' || $1 || '%'
ORDER BY email
    LIMIT $2 OFFSET $3;
