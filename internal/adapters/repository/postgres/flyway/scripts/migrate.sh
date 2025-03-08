#!/usr/bin/env bash

# Extract database connection details from DATABASE_URL
# Format: postgres://username:password@host:port/database
if [[ $DATABASE_URL =~ postgres://([^:]+):([^@]+)@([^:]+):([0-9]+)/(.+) ]]; then
  DB_USER="${BASH_REMATCH[1]}"
  DB_PASS="${BASH_REMATCH[2]}"
  DB_HOST="${BASH_REMATCH[3]}"
  DB_PORT="${BASH_REMATCH[4]}"
  DB_NAME="${BASH_REMATCH[5]}"
else
  echo "Error: Invalid DATABASE_URL format"
  exit 1
fi

export FLYWAY_USER="${DB_USER}"
export FLYWAY_PASSWORD="${DB_PASS}"
export FLYWAY_URL="jdbc:postgresql://${DB_HOST}:${DB_PORT}/${DB_NAME}"

flyway -connectRetries=5 -skipCheckForUpdate -validateMigrationNaming=true -locations=migrations -schemas=flyway migrate
