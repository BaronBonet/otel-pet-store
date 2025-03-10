#!/bin/bash

export PATH="$(pwd)/local/bin:$PATH"
echo "✓ Added local/bin to PATH"

set -a
source .env
set +a
echo "✓ Loaded environment variables from .env"
