package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pipat557/golang-graphql/graph/generated"
	"github.com/pipat557/golang-graphql/graph/model"
)

func (r *mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (*model.Owner, error) {
	owner := model.Owner{
		Name:     input.Name,
		IsSingle: input.IsSingle,
		// Address:  input.Address,
	}

	_, err := r.DB.Model(&owner).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new Owner: %v", err)
	}

	return &owner, nil
}

func (r *mutationResolver) UpdateOwner(ctx context.Context, input model.OldOwner) (*model.Owner, error) {
	owner := model.Owner{
		ID:       input.ID,
		Name:     input.Name,
		IsSingle: input.IsSingle,
	}

	_, err := r.DB.Model(&owner).WherePK().UpdateNotNull()
	if err != nil {
		return nil, fmt.Errorf("error updating Owner: %v", err)
	}

	return &owner, nil
}

func (r *mutationResolver) DeleteOwner(ctx context.Context, input model.Delete) (*model.Owner, error) {
	owner := model.Owner{
		ID: input.ID,
	}

	_, err := r.DB.Model(&owner).WherePK().Delete()
	if err != nil {
		return nil, fmt.Errorf("error updating Owner: %v", err)
	}

	return &owner, nil
}

func (r *queryResolver) Owner(ctx context.Context) ([]*model.Owner, error) {
	var owner []*model.Owner

	err := r.DB.Model(&owner).Select()
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (r *queryResolver) SingleOwner(ctx context.Context, isSingle bool) ([]*model.Owner, error) {
	var owner []*model.Owner

	err := r.DB.Model(&owner).Where("is_single=?", isSingle).Select()
	if err != nil {
		return nil, err
	}

	return owner, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
