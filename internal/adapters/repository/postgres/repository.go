package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/BaronBonet/otel-pet-store/internal/adapters/repository/postgres/generated"
	"github.com/BaronBonet/otel-pet-store/internal/core"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	// Convert string UUID to pgtype.UUID
	id, err := uuid.Parse(pet.ID)
	if err != nil {
		return fmt.Errorf("parsing pet ID: %w", err)
	}
	pgUUID := pgtype.UUID{Bytes: id, Valid: true}

	// Convert time.Time to pgtype.Timestamptz
	createdAt := pgtype.Timestamptz{Time: pet.CreatedAt, Valid: true}
	updatedAt := pgtype.Timestamptz{Time: pet.UpdatedAt, Valid: true}

	err = r.queries.CreatePet(ctx, generated.CreatePetParams{
		ID:        pgUUID,
		Name:      pet.Name,
		Type:      string(pet.Type),
		Status:    string(pet.Status),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return fmt.Errorf("creating pet in database: %w", err)
	}
	return nil
}

func (r *repository) GetPet(ctx context.Context, id string) (*core.Pet, error) {
	// Convert string UUID to pgtype.UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("parsing pet ID: %w", err)
	}
	pgUUID := pgtype.UUID{Bytes: uid, Valid: true}

	pet, err := r.queries.GetPet(ctx, pgUUID)
	if err != nil {
		return nil, fmt.Errorf("getting pet from database: %w", err)
	}

	return &core.Pet{
		ID:        uuid.UUID(pet.ID.Bytes).String(),
		Name:      pet.Name,
		Type:      core.PetType(pet.Type),
		Status:    core.PetStatus(pet.Status),
		CreatedAt: pet.CreatedAt.Time,
		UpdatedAt: pet.UpdatedAt.Time,
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
			ID:        uuid.UUID(pet.ID.Bytes).String(),
			Name:      pet.Name,
			Type:      core.PetType(pet.Type),
			Status:    core.PetStatus(pet.Status),
			CreatedAt: pet.CreatedAt.Time,
			UpdatedAt: pet.UpdatedAt.Time,
		}
	}
	return result, nil
}

func (r *repository) UpdatePetStatus(ctx context.Context, id string, status core.PetStatus) error {
	// Convert string UUID to pgtype.UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("parsing pet ID: %w", err)
	}
	pgUUID := pgtype.UUID{Bytes: uid, Valid: true}

	// Convert time.Time to pgtype.Timestamptz
	updatedAt := pgtype.Timestamptz{Time: time.Now(), Valid: true}

	err = r.queries.UpdatePetStatus(ctx, generated.UpdatePetStatusParams{
		ID:        pgUUID,
		Status:    string(status),
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return fmt.Errorf("updating pet status in database: %w", err)
	}
	return nil
}
