package user

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/amaru0601/go-rent/ent"
	"github.com/amaru0601/go-rent/ent/contract"
	"github.com/amaru0601/go-rent/ent/predicate"
	"github.com/amaru0601/go-rent/ent/property"
	"github.com/amaru0601/go-rent/ent/user"
)

func (r *contractResolver) Tenant(ctx context.Context, obj *ent.Contract) (*ent.User, error) {
	//TODO improve query to get Tenant
	return r.client.User.Query().Where(
		user.HasContractsWith(
			predicate.Contract(user.HasPropertiesWith(
				property.HasContractWith(
					contract.IDEQ(obj.ID),
				),
			)),
		),
	).Only(ctx)
}

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
		SetAddress(property.Address).
		SetCity(property.City).
		SetDescription(property.Description).
		SetOwnerID(property.OwnerID).Save(ctx)
}

func (r *mutationResolver) CreateContract(ctx context.Context, contract ContractInput) (*ent.Contract, error) {
	contract_, err := r.client.Contract.Create().
		SetStartDate(contract.StartDate).
		SetEndDate(contract.EndDate).
		SetPayAmount(contract.PayAmount).
		SetPayDate(contract.PayDate).
		SetPropertyID(contract.PropertyID).Save(ctx)

	owner, err := r.client.User.Query().Where(user.IDEQ(contract.OwnerID)).Only(ctx)
	err = owner.Update().AddContractIDs(contract_.ID).Exec(ctx)

	tenant, err := r.client.User.Query().Where(user.IDEQ(contract.TenantID)).Only(ctx)
	err = tenant.Update().AddContractIDs(contract_.ID).Exec(ctx)

	//TODO: UPDATE PROPERTY CONTRACT ID

	return contract_, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Contract returns ContractResolver implementation.
func (r *Resolver) Contract() ContractResolver { return &contractResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type contractResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
