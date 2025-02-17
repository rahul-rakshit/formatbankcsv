#!/bin/bash

set -euo pipefail

# Define variables
VERSION="v1.2.0"
OUTPUT_DIR="dist"
BINARY_NAME="formatbankcsv"
PLATFORMS=("darwin/arm64" "darwin/amd64" "linux/arm64" "linux/amd64")

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Build and package for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
    OS=${PLATFORM%/*}
    ARCH=${PLATFORM#*/}
    ZIP_NAME="${OS}-${ARCH}-${VERSION}.zip"
    OUTPUT_BINARY="${OUTPUT_DIR}/${BINARY_NAME}"

    echo "Building for ${OS}/${ARCH}..."
    GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT_BINARY" .

    echo "Packaging ${ZIP_NAME}..."
    zip -j "${OUTPUT_DIR}/${ZIP_NAME}" "$OUTPUT_BINARY"

    # Remove the binary after zipping
    rm "$OUTPUT_BINARY"
done

echo "Build and packaging completed. Files are in the '${OUTPUT_DIR}' directory."
