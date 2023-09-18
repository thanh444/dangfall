package animal

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/animal"
	"golang-boilerplate/model"
	"strings"

	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, input model.CreateAnimalInput) (*ent.Animal, error)
	List(ctx context.Context, filter ent.AnimalFilterInput) (*ent.AnimalConnection, error)
}

type impla struct {
	client *ent.Client
}

func New(client *ent.Client) Repository {
	return &impla{
		client: client,
	}
}

// Create implements Repository.
func (a *impla) Create(ctx context.Context, input model.CreateAnimalInput) (*ent.Animal, error) {
	animal, err := a.client.Animal.Create().SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return animal, nil
}

// List implements Repository.
func (a *impla) List(ctx context.Context, filter ent.AnimalFilterInput) (*ent.AnimalConnection, error) {
	query := a.client.Animal.Query()
	if filter.Name != nil {
		query = query.Where((animal.NameContainsFold(strings.TrimSpace(*filter.Name))))
	}
	animal, err := query.Paginate(ctx, filter.Pagination.After, filter.Pagination.Last, filter.Pagination.Before, filter.Pagination.First)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return animal, nil
}

// tao mot Animal  moi
