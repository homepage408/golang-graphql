package handler

import (
	"context"

	"github.com/golang-graphql/usecase"
)

type Handler interface {
	CreateUser(ctx context.Context)
	GetDetailUser(ctx context.Context)
}

type handler struct {
	usecase usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *handler {
	return &handler{
		usecase: uc,
	}
}
