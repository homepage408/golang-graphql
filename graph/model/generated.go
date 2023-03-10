// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type UserResponseOk interface {
	IsUserResponseOk()
}

type Error struct {
	Message *string `json:"message"`
	Success *bool   `json:"success"`
}

func (Error) IsUserResponseOk() {}

type GetUserParam struct {
	ID *int `json:"ID"`
}

type UserResponse struct {
	ID     *int    `json:"ID"`
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Alamat *string `json:"alamat"`
}

func (UserResponse) IsUserResponseOk() {}
