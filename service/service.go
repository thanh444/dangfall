package service

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/repository"
	"golang-boilerplate/service/animal"
	"golang-boilerplate/service/user"

	"go.uber.org/zap"
)

// ServiceRegistry is the interface for the service registry.
type ServiceRegistry interface {
	User() user.Service
	Animal() animal.Service
}

// impl is the implementation of the service registry.
type impl struct {
	user   user.Service
	animal animal.Service
}

// Animal implements ServiceRegistry.

// New creates a new service registry.
func New(entClient *ent.Client, logger *zap.Logger) ServiceRegistry {
	repoRegistry := repository.New(entClient)

	return &impl{
		user: user.New(repoRegistry, logger),
	}
}

func New1(entClient *ent.Client, logger *zap.Logger) ServiceRegistry {
	repoRegistry := repository.New(entClient)

	return &impl{
		animal: animal.New(repoRegistry, logger),
	}
}

// User returns the user service.
func (i impl) User() user.Service {
	return i.user
}
func (*impl) Animal() animal.Service {
	panic("unimplemented")
}
