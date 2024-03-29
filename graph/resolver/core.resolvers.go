package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/golang-graphql/graph"
	"github.com/golang-graphql/graph/model"
)

// User is the resolver for the User field.
func (r *mutationResolver) User(ctx context.Context) (*model.AbstracModel, error) {
	// panic(fmt.Errorf("not implemented: User - User"))
	return new(model.AbstracModel), nil
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context) (*model.AbstracModel, error) {
	// panic(fmt.Errorf("not implemented: User - User"))
	return new(model.AbstracModel), nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
