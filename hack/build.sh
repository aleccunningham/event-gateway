#!/bin/sh
set -e
[ -z "$DIST" ] && DIST=.bin

[ -z "$VERSION" ] && VERSION=$(cat VERSION)
[ -z "$SHORT_SHA" ] && VCS_COMMIT_ID=$(git rev-parse --short HEAD 2>/dev/null)

pkg="github.com/marjoram/api-glue/pkg"

echo "VERSION: $VERSION"
echo "SHORT_SHA: $SHORT_SHA"

go_build() {
  [ -d "${DIST}" ] && rm -rf "${DIST:?}/*"
  [ -d "${DIST}" ] || mkdir -p "${DIST}"
  CGO_ENABLED=0 go build \
    -ldflags "-s -w -X $pkg/version.SemVer=${VERSION} -X $pkg/version.GitCommit=${SHORT_SHA}" \
    -o "${DIST}/hermes" ./cmd
}

go_build
