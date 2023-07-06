package usecase

import (
	"context"

	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/repository"
)

type Usecase interface {
	CreateUser(ctx context.Context, input *model.SignInParams) (*model.MainResponse, error)
	GetDetailUser(ctx context.Context, input *model.GetUserParam) (*model.MainResponseList, error)
}

type usecase struct {
	repository repository.Repository
	saltText   string
}

func NewUsecase(repository repository.Repository, saltText string) *usecase {
	return &usecase{
		repository: repository,
		saltText:   saltText,
	}
}
