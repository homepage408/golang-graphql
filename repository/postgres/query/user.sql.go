// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package query

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
Insert into users (name, email, password, address, updated_at) values($1,$2,$3,$4, now()) RETURNING id ,name, email, address
`

type CreateUserParams struct {
	Name     sql.NullString
	Email    sql.NullString
	Password sql.NullString
	Address  sql.NullString
}

type CreateUserRow struct {
	ID      int32
	Name    sql.NullString
	Email   sql.NullString
	Address sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Address,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Address,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, name, email from users where email = $1
`

type GetUserByEmailRow struct {
	ID    int32
	Name  sql.NullString
	Email sql.NullString
}

func (q *Queries) GetUserByEmail(ctx context.Context, email sql.NullString) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const getUserDetail = `-- name: GetUserDetail :one
select id, name, email, address from users where id = $1
`

type GetUserDetailRow struct {
	ID      int32
	Name    sql.NullString
	Email   sql.NullString
	Address sql.NullString
}

func (q *Queries) GetUserDetail(ctx context.Context, id int32) (GetUserDetailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserDetail, id)
	var i GetUserDetailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Address,
	)
	return i, err
}
