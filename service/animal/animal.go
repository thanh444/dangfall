package animal

import (
	"context"
	"errors"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
	"golang-boilerplate/repository"

	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, input model.CreateAnimalInput) (*ent.Animal, error)
	List(ctx context.Context, filter ent.AnimalFilterInput) (*ent.AnimalConnection, error)
}

type impl struct {
	repoRegistry repository.RepositoryRegistry
	logger       *zap.Logger
}

// Create implements Service.
func (a *impl) Create(ctx context.Context, input model.CreateAnimalInput) (*ent.Animal, error) {
	animal, err := a.repoRegistry.Animal().Create(ctx, input)
	if err != nil {
		a.logger.Error("create_animal_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "create_animal_failed")
	}
	return animal, nil
}

// List implements Service.
func (a *impl) List(ctx context.Context, filter ent.AnimalFilterInput) (*ent.AnimalConnection, error) {
	animals, err := a.repoRegistry.Animal().List(ctx, filter)
	if err != nil {
		a.logger.Error("list_animal_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "list_animal_has_failed")
	}
	return animals, nil
}

func New(repoRegistry repository.RepositoryRegistry, logger *zap.Logger) Service {
	return &impl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}
