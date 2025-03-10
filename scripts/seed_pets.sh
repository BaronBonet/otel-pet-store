#!/bin/bash

# Function to create a pet using grpcurl
create_pet() {
    local name=$1
    local type=$2
    grpcurl -d "{\"name\": \"$name\", \"type\": \"$type\"}" \
        -proto api/petstore/v1/rpc.proto \
        -import-path api \
        -plaintext \
        localhost:8080 \
        petstore.v1.PetStoreService/CreatePet
}

echo "Creating sample pets..."

# Create dogs
create_pet "Max" "PET_TYPE_DOG"
create_pet "Luna" "PET_TYPE_DOG"
create_pet "Rocky" "PET_TYPE_DOG"

# Create cats
create_pet "Milo" "PET_TYPE_CAT"
create_pet "Lucy" "PET_TYPE_CAT"
create_pet "Oliver" "PET_TYPE_CAT"

# Create rabbits
create_pet "Thumper" "PET_TYPE_RABBIT"
create_pet "Bunny" "PET_TYPE_RABBIT"
create_pet "Hoppy" "PET_TYPE_RABBIT"

echo "Done creating sample pets!"

# List all pets to verify
echo -e "\nListing all pets:"
grpcurl -proto api/petstore/v1/rpc.proto \
    -import-path api \
    -plaintext \
    localhost:8080 \
    petstore.v1.PetStoreService/ListPets
