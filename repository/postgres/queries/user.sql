-- name: GetUserDetail :one
select id, name, email, address from users where id = $1;

-- name: CreateUser :one
Insert into users (name, email, password, address, updated_at) values($1,$2,$3,$4, now()) RETURNING id ,name, email, address;

-- name: GetUserByEmail :one
select id, name, email from users where email = $1;