package core

import (
	"context"
	"fmt"
	"time"

	"github.com/BaronBonet/otel-pet-store/internal/pkg/logger"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type service struct {
	repo        Repository
	logger      logger.Logger
	meter       metric.Meter
	petsCreated metric.Int64Counter
}

func NewService(repo Repository, logger logger.Logger) Service {
	meter := otel.Meter("github.com/BaronBonet/otel-pet-store")
	petsCreated, err := meter.Int64Counter(
		"pets_created",
		metric.WithDescription("Counts the number of pets created"),
	)
	if err != nil {
		logger.Error(context.Background(), "Failed to create counter", "error", err)
	}
	return &service{repo: repo, logger: logger, meter: meter, petsCreated: petsCreated}
}

func (s *service) CreatePet(ctx context.Context, name string, petType PetType) (*Pet, error) {
	s.logger.Info(ctx, "Creating pet", "name", name, "type", petType)
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

	s.petsCreated.Add(ctx, 1)

	return pet, nil
}

func (s *service) GetPet(ctx context.Context, id string) (*Pet, error) {
	s.logger.Info(ctx, "Getting pet", "id", id)
	return s.repo.GetPet(ctx, id)
}

func (s *service) ListPets(ctx context.Context) ([]*Pet, error) {
	s.logger.Info(ctx, "Listing pets")
	return s.repo.ListPets(ctx)
}

func (s *service) UpdatePetStatus(ctx context.Context, id string, status PetStatus) (*Pet, error) {
	s.logger.Info(ctx, "Updating pet status", "id", id, "status", status)
	if err := s.repo.UpdatePetStatus(ctx, id, status); err != nil {
		return nil, fmt.Errorf("updating pet status: %w", err)
	}
	return s.repo.GetPet(ctx, id)
}
