package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"
	"golang-boilerplate/ent"
	graphql1 "golang-boilerplate/graphql"
	"golang-boilerplate/model"
)

// ID is the resolver for the id field.
func (r *animalResolver) ID(ctx context.Context, obj *ent.Animal) (string, error) {
	return obj.ID.String(), nil
}

// Create is the resolver for the create field.
func (r *animalOpsResolver) Create(ctx context.Context, obj *ent.AnimalOps, input model.CreateAnimalInput) (*ent.Animal, error) {
	return r.service.Animal().Create(ctx, input)
}

// List is the resolver for the list field.
func (r *animalQueriesResolver) List(ctx context.Context, obj *model.AnimalQueries, filter *ent.AnimalFilterInput) (*ent.AnimalConnection, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}

// Animal is the resolver for the animal field.
func (r *mutationResolver) Animal(ctx context.Context) (*ent.AnimalOps, error) {
	return &ent.AnimalOps{}, nil
}

// Animal is the resolver for the animal field.
func (r *queryResolver) Animal(ctx context.Context) (*model.AnimalQueries, error) {
	return &model.AnimalQueries{}, nil
}

// Order is the resolver for the order field.
func (r *animalFilterInputResolver) Order(ctx context.Context, obj *ent.AnimalFilterInput, data *ent.AnimalOrder) error {
	panic(fmt.Errorf("not implemented: Order - order"))
}

// Animal returns graphql1.AnimalResolver implementation.
func (r *Resolver) Animal() graphql1.AnimalResolver { return &animalResolver{r} }

// AnimalOps returns graphql1.AnimalOpsResolver implementation.
func (r *Resolver) AnimalOps() graphql1.AnimalOpsResolver { return &animalOpsResolver{r} }

// AnimalQueries returns graphql1.AnimalQueriesResolver implementation.
func (r *Resolver) AnimalQueries() graphql1.AnimalQueriesResolver { return &animalQueriesResolver{r} }

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

// AnimalFilterInput returns graphql1.AnimalFilterInputResolver implementation.
func (r *Resolver) AnimalFilterInput() graphql1.AnimalFilterInputResolver {
	return &animalFilterInputResolver{r}
}

type animalResolver struct{ *Resolver }
type animalOpsResolver struct{ *Resolver }
type animalQueriesResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type animalFilterInputResolver struct{ *Resolver }
