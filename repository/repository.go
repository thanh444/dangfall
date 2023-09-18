package repository

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/repository/animal"
	"golang-boilerplate/repository/user"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	User() user.Repository
	Animal() animal.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	user1  user.Repository
	animal animal.Repository
}

// Animal implements RepositoryRegistry.

// New creates a new repository registry.
func New(client *ent.Client) RepositoryRegistry {
	return &impl{
		user1: user.New(client),
	}
}
func NewA(client *ent.Client) RepositoryRegistry {
	return &impl{
		animal: animal.New(client),
	}
}

// User returns the user repository.
func (i impl) User() user.Repository {
	return i.user1
}
func (i *impl) Animal() animal.Repository {
	return i.animal
}
