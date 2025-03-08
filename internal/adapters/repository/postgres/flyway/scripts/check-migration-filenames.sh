#!/usr/bin/env bash

# Ensure filenames conform with specifications from flyway
# https://documentation.red-gate.com/fd/migrations-184127470.html
# - Start with either V,U or R
# - Are sequential
# - Any U migration must have corresponding V migration filename

unknownName=""
lastV=0
declare -A v_files

# Create an array of file names that start with a V
for f in "$@"; do
  file=$(basename "$f")
  if [[ "$file" =~ ^V[0-9]{7}__ ]]; then
    number=$(echo "$file" | sed 's/^V\([0-9]\{7\}\)__.*$/\1/')
    v_files[$number]=1
  fi
done

for f in "$@"; do
  file=$(basename "$f")

  # Handle R files separately as they don't have numbers
  if [[ "$file" =~ ^R__[A-Za-z0-9_]+\.sql$ ]]; then
    continue
  fi

  # Check for correct format: V/U followed by exactly 7 digits, double underscore, name, and .sql extension
  if [[ ! "$file" =~ ^[VU][0-9]{7}__[A-Za-z0-9_]+\.sql$ ]]; then
    echo "Error: File $file does not match required format (e.g., V0000001__NAME_OF_MIGRATION.sql)"
    unknownName=$file
    continue
  fi

  # Extract the type (V/U) and number
  type=${file:0:1}
  number=$(echo "$file" | sed 's/^[VU]\([0-9]\{7\}\)__.*$/\1/')
  number=$((10#$number)) # Force base-10 interpretation

  if [[ $type == "V" ]]; then
    if [[ $lastV -eq 0 ]]; then
      # First V file must start with 1
      if [[ $number -ne 1 ]]; then
        echo "Error: First version file must start with V0000001, got $file"
        unknownName=$file
        continue
      fi
    elif [[ $number -ne $((lastV + 1)) ]]; then
      echo "Error: File $file is not sequential. Expected V$(printf "%07d" $((lastV + 1))), skipping numbers is not allowed"
      unknownName=$file
      continue
    fi
    lastV=$number
  else
    # For U files, check if matching V file exists
    if [[ -z "${v_files[$(printf "%07d" $number)]}" ]]; then
      echo "Error: Undo file $file has no matching version file"
      unknownName=$file
      continue
    fi
  fi
done

# Check for any files that don't match our expected patterns
for f in "$@"; do
  file=$(basename "$f")
  if [[ ! "$file" =~ ^V[0-9]{7}__[A-Za-z0-9_]+\.sql$ ]] &&
    [[ ! "$file" =~ ^U[0-9]{7}__[A-Za-z0-9_]+\.sql$ ]] &&
    [[ ! "$file" =~ ^R__[A-Za-z0-9_]+\.sql$ ]]; then
    echo "Error: File $file does not match any valid pattern (V######__name.sql, U######__name.sql, or R__name.sql)"
    unknownName=$file
  fi
done

if [[ -n "$unknownName" ]]; then
  exit 1
fi
