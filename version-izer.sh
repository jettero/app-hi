#!/usr/bin/env bash
#
# https://belief-driven-design.com/build-time-variables-in-go-51439b26ef9/
# (broken, but fixed via)
# https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications

# STEP 1: Determinate the required values

VERSION="$(git describe --tags --always --dirty --match='v[0-9].*' | sed -e 's/-g\([a-f0-9]\)/-\1/')"
COMMIT_HASH="$(git rev-parse HEAD)"
BUILD_TIMESTAMP="$(date -u)"

# STEP 2: Build the ldflags

LDFLAGS=(
  "-X 'main.Version=${VERSION}'"
  "-X 'main.CommitHash=${COMMIT_HASH}'"
  "-X 'main.BuildTime=${BUILD_TIMESTAMP}'"
)

# STEP 3: Actual Go build process

set -x

go build -ldflags="${LDFLAGS[*]}" "$@"
