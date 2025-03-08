package core

import "time"

type Pet struct {
	ID        string
	Name      string
	Type      PetType
	Status    PetStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PetType string

const (
	PetTypeDog    PetType = "DOG"
	PetTypeCat    PetType = "CAT"
	PetTypeRabbit PetType = "RABBIT"
)

type PetStatus string

const (
	PetStatusAvailable PetStatus = "AVAILABLE"
	PetStatusPending   PetStatus = "PENDING"
	PetStatusSold      PetStatus = "SOLD"
)
