package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/BaronBonet/otel-pet-store/internal/core"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	queries *generated.Queries
}

func NewRepository(pool *pgxpool.Pool) core.Repository {
	return &repository{
		queries: generated.New(pool),
	}
}

func (r *repository) CreatePet(ctx context.Context, pet *core.Pet) error {
	err := r.queries.CreatePet(ctx, generated.CreatePetParams{
		ID:        pet.ID,
		Name:      pet.Name,
		Type:      generated.PetType(pet.Type),
		Status:    generated.PetStatus(pet.Status),
		CreatedAt: pet.CreatedAt,
		UpdatedAt: pet.UpdatedAt,
	})
	if err != nil {
		return fmt.Errorf("creating pet in database: %w", err)
	}
	return nil
}

func (r *repository) GetPet(ctx context.Context, id string) (*core.Pet, error) {
	pet, err := r.queries.GetPet(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting pet from database: %w", err)
	}
	return &core.Pet{
		ID:        pet.ID,
		Name:      pet.Name,
		Type:      core.PetType(pet.Type),
		Status:    core.PetStatus(pet.Status),
		CreatedAt: pet.CreatedAt,
		UpdatedAt: pet.UpdatedAt,
	}, nil
}

func (r *repository) ListPets(ctx context.Context) ([]*core.Pet, error) {
	pets, err := r.queries.ListPets(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing pets from database: %w", err)
	}
	result := make([]*core.Pet, len(pets))
	for i, pet := range pets {
		result[i] = &core.Pet{
			ID:        pet.ID,
			Name:      pet.Name,
			Type:      core.PetType(pet.Type),
			Status:    core.PetStatus(pet.Status),
			CreatedAt: pet.CreatedAt,
			UpdatedAt: pet.UpdatedAt,
		}
	}
	return result, nil
}

func (r *repository) UpdatePetStatus(ctx context.Context, id string, status core.PetStatus) error {
	err := r.queries.UpdatePetStatus(ctx, generated.UpdatePetStatusParams{
		ID:        id,
		Status:    generated.PetStatus(status),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("updating pet status in database: %w", err)
	}
	return nil
}
