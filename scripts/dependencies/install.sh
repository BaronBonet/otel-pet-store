#!/usr/bin/env bash
set -euo pipefail

DIR_PATH=$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)

. "${DIR_PATH}/../../build_dependencies_versions"

PLATFORM=$(uname)
ARCH=$(uname -m)

PLATFORM_LOWER=$(echo "$PLATFORM" | awk '{print tolower($0)}')
ARCH_LOWER=$(echo "$ARCH" | awk '{print tolower($0)}')

# Determine OS based on platform
if [[ "$PLATFORM_LOWER" == "darwin" ]]; then
  OS="macos"
elif [[ "$PLATFORM_LOWER" == "linux" ]]; then
  OS="linux"
else
  echo "Unsupported platform: $PLATFORM"
  exit 1
fi

DEST_DIR="${DIR_PATH}/../../local/bin"

echo "Downloading sqlc $SQLC_VERSION"
curl -s -o "${DEST_DIR}/sqlc.tar.gz" \
  -L "https://github.com/kyleconroy/sqlc/releases/download/v${SQLC_VERSION}/sqlc_${SQLC_VERSION}_${PLATFORM_LOWER}_${ARCH_LOWER}.tar.gz"
echo "Extracting sqlc to $DEST_DIR"
tar -xf "${DEST_DIR}/sqlc.tar.gz" -C "${DEST_DIR}/"
echo "Cleaning up sqlc archive"
rm "${DEST_DIR}/sqlc.tar.gz"

echo "Downloading sqlfluff"
pipx install --index-url https://pypi.org/simple sqlfluff=="$SQLFLUFF_VERSION"

echo "installing golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "${DEST_DIR}" v${GOLANGCI_LINT_VERSION}
