-- name: GetUserDetail :one
select id, name, email, address from users where id = $1;

-- name: CreateUser :one
Insert into users (name, email, password, address) values($1,$2,$3,$4) RETURNING name, email, address;