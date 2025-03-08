package core

import "context"

type Service interface {
	CreatePet(ctx context.Context, name string, petType PetType) (*Pet, error)
	GetPet(ctx context.Context, id string) (*Pet, error)
	ListPets(ctx context.Context) ([]*Pet, error)
	UpdatePetStatus(ctx context.Context, id string, status PetStatus) (*Pet, error)
}

type Repository interface {
	CreatePet(ctx context.Context, pet *Pet) error
	GetPet(ctx context.Context, id string) (*Pet, error)
	ListPets(ctx context.Context) ([]*Pet, error)
	UpdatePetStatus(ctx context.Context, id string, status PetStatus) error
}
