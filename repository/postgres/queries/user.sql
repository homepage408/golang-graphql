-- name: GetUserDetail :one
select id, name, email, alamat from users where id = $1;

-- name: CreateUser :one
Insert into users (id, name, email, alamat) values($1,$2,$3,$4) RETURNING id;