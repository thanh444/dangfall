// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"golang-boilerplate/ent"
)

type AnimalQueries struct {
	List *ent.AnimalConnection `json:"list"`
}

type CreateAnimalInput struct {
	Name string `json:"name"`
}

type CreateUserInput struct {
	Name string `json:"name"`
}

type UserQueries struct {
	List *ent.UserConnection `json:"list"`
}