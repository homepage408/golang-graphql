// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package query

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
Insert into users (id, name, email, alamat) values($1,$2,$3,$4) RETURNING id
`

type CreateUserParams struct {
	ID     int32
	Name   sql.NullString
	Email  sql.NullString
	Alamat sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Alamat,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserDetail = `-- name: GetUserDetail :one
select id, name, email, alamat from users where id = $1
`

type GetUserDetailRow struct {
	ID     int32
	Name   sql.NullString
	Email  sql.NullString
	Alamat sql.NullString
}

func (q *Queries) GetUserDetail(ctx context.Context, id int32) (GetUserDetailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserDetail, id)
	var i GetUserDetailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Alamat,
	)
	return i, err
}
