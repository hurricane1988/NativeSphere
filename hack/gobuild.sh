set -o errexit
set -o nounset
set -o pipefail

# Output path of binary.
OUTPUT_DIR=bin
# Go env GOOS.
BUILD_GOOS=linux
# Setup architecture.
BUILD_GOARCH=amd64
# Disable the CGG(CGO can't been supported by cross compile)
ENABLE_CGO=0
# Setup the version of compile.
LDFLAGS='-X '

# forgoing -i (incremental build) because it will be deprecated by tool chain.
GOOS=${BUILD_GOOS} CGO_ENABLED=${ENABLE_CGO} GOARCH=${BUILD_GOARCH} ${GOBINARY} build \
        -ldflags="${LDFLAGS}" \
        -o "${OUT}" \
        "${BUILDPATH}"