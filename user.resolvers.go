package user

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/amaru0601/go-rent/ent"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user UserInput) (*ent.User, error) {
	return r.client.User.Create().
		SetNames(user.Names).
		SetLastnames(user.Lastnames).
		SetEmail(user.Email).
		SetActivate(true).
		SetCreatedAt(time.Now()).
		SetBirthday(*user.Birthday).Save(ctx)
}

func (r *mutationResolver) CreateProperty(ctx context.Context, property PropertyInput) (*ent.Property, error) {
	return r.client.Property.Create().
		SetClass(property.Class).
		SetDeleted(false).
		SetName(property.Name).
		SetOwnerID(property.Owner).Save(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
