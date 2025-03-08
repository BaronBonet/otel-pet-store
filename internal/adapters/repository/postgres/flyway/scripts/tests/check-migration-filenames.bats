#!/usr/bin/env bats

setup() {
  # Create a temporary directory for test files
  TEST_DIR="$(mktemp -d)"
  # Get the absolute path to the script being tested
  SCRIPT_PATH="$(cd "$(dirname "$BATS_TEST_FILENAME")/.." && pwd)/check-migration-filenames.sh"
}

teardown() {
  # Clean up temporary directory after each test
  rm -rf "$TEST_DIR"
}

# Helper function to create test files
create_test_files() {
  for filename in "$@"; do
    touch "$TEST_DIR/$filename"
  done
}

@test "valid V files sequence passes" {
  create_test_files \
    "V0000001__First.sql" \
    "V0000002__Second.sql" \
    "V0000003__Third.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 0 ]
}

@test "valid V and U files sequence passes" {
  create_test_files \
    "V0000001__First.sql" \
    "V0000002__Second.sql" \
    "U0000001__Undo_First.sql" \
    "U0000002__Undo_Second.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 0 ]
}

@test "valid R files pass" {
  create_test_files \
    "R__View.sql" \
    "R__MaterializedView.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 0 ]
}

@test "fails when V files skip numbers" {
  create_test_files \
    "V0000001__First.sql" \
    "V0000003__Third.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" == *"not sequential"* ]]
}

@test "fails when V doesn't start with 0000001" {
  create_test_files "V0000002__First.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" == *"must start with V0000001"* ]]
}

@test "fails when U file has no matching V file" {
  create_test_files \
    "V0000001__First.sql" \
    "U0000002__Second.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" == *"no matching version file"* ]]
}

@test "fails with invalid format" {
  create_test_files "V1__Invalid.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" == *"does not match required format"* ]]
}

@test "R files with numbers should fail" {
  create_test_files "R0000001__Invalid.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" == *"does not match required format"* ]]
}

@test "mixed valid files pass" {
  create_test_files \
    "V0000001__First.sql" \
    "V0000002__Second.sql" \
    "U0000001__Undo_First.sql" \
    "R__View.sql" \
    "R__Another_View.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 0 ]
}

@test "fails if there is a name that does not match our expected pattern" {
  create_test_files \
    "V0000001__First.sql" \
    "BAD_FILE_NAME.sql"

  run "$SCRIPT_PATH" "$TEST_DIR"/*
  [ "$status" -eq 1 ]
}
