package handler

import (
	"context"

	"github.com/golang-graphql/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (h *handler) CreateUser(ctx context.Context, param *model.SignInParams) (*model.MainResponse, error) {
	// inputParam :=
	res, err := h.usecase.CreateUser(ctx, param)

	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}
	return res, nil

}

func (h *handler) GetDetailUser(ctx context.Context, param *model.GetUserParam) (*model.MainResponseList, error) {
	return nil, nil
}
