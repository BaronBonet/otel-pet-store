package core

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePet(ctx context.Context, name string, petType PetType) (*Pet, error) {
	pet := &Pet{
		ID:        uuid.New().String(),
		Name:      name,
		Type:      petType,
		Status:    PetStatusAvailable,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreatePet(ctx, pet); err != nil {
		return nil, fmt.Errorf("creating pet: %w", err)
	}

	return pet, nil
}

func (s *service) GetPet(ctx context.Context, id string) (*Pet, error) {
	return s.repo.GetPet(ctx, id)
}

func (s *service) ListPets(ctx context.Context) ([]*Pet, error) {
	return s.repo.ListPets(ctx)
}

func (s *service) UpdatePetStatus(ctx context.Context, id string, status PetStatus) (*Pet, error) {
	if err := s.repo.UpdatePetStatus(ctx, id, status); err != nil {
		return nil, fmt.Errorf("updating pet status: %w", err)
	}
	return s.repo.GetPet(ctx, id)
}
