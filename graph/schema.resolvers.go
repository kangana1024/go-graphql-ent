package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kangana1024/go-graphql-ent/ent"
	"github.com/kangana1024/go-graphql-ent/graph/generated"
	"github.com/kangana1024/go-graphql-ent/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// BirthDate is the resolver for the birthDate field.
func (r *userResolver) BirthDate(ctx context.Context, obj *ent.User) (*string, error) {
	panic(fmt.Errorf("not implemented: BirthDate - birthDate"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (*string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *ent.User) (*string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *ent.User) (*string, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
