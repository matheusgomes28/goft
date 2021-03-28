#!/usr/bin/env bash

set -euo pipefail

: ' This script will find all the "*.go" files and
    run the "gofmt -d -s" command on them. Note that
    the "gofmt -d -s" command will fail if it thinks
    that a file needs to be formatted, and it will
    not change the files themselves!
'

REPOSITORY_ROOT="$(cd "../" && pwd)"
GO_FILES="$(find "$REPOSITORY_ROOT" -name "*.go")"

# Test if the gofmt prints empty string, meaning no
# changes are needed to any go file
# shellcheck disable=SC2086
LINT_OUTPUT="$(gofmt -d -s $GO_FILES)"
echo "$LINT_OUTPUT"
test -z "$LINT_OUTPUT"
