#!/usr/bin/env bash

exec go run ${DFMT:+-tags debug} ./cmd/hi "$@"
