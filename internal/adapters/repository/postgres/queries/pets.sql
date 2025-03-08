-- name: CreatePet :exec
INSERT INTO pets (
    id, name, type, status, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetPet :one
SELECT * FROM pets WHERE id = $1;

-- name: ListPets :many
SELECT * FROM pets;

-- name: UpdatePetStatus :exec
UPDATE pets 
SET status = $2, updated_at = $3
WHERE id = $1; 