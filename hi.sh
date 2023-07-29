#!/usr/bin/env bash

exec go run -tags debug ./cmd/hi "$@"
