#!/usr/bin/env bash
set -euo pipefail

DIR_PATH=$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)
. "${DIR_PATH}/../../build_dependencies_versions"

LOCAL_BIN="${DIR_PATH}/../../local/bin"

exit_code=0
if [[ ! -f "$LOCAL_BIN/sqlc" ]]; then
  echo "sqlc is not installed."
  exit_code=1
fi

if [[ $exit_code != 0 ]]; then
  exit ${exit_code}
fi

ACTUAL_SQLC_VERSION="$("${LOCAL_BIN}/sqlc" version)"

if [[ "${ACTUAL_SQLC_VERSION}" != "v${SQLC_VERSION}" ]]; then
  echo "sqlc version ($ACTUAL_SQLC_VERSION) mismatch - expected ${SQLC_VERSION}"
  exit_code=1
fi

ACTUAL_GOLANGCI_LINT_VERSION=$("${LOCAL_BIN}/golangci-lint" version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)

if [[ "${ACTUAL_GOLANGCI_LINT_VERSION}" != "${GOLANGCI_LINT_VERSION}" ]]; then
  echo "golangci-lint version (${ACTUAL_GOLANGCI_LINT_VERSION}) mismatch - expected ${GOLANGCI_LINT_VERSION}"
  exit_code=1
fi

exit ${exit_code}
