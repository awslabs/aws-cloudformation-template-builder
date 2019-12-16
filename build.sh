#!/bin/bash

# This script will build cfn-skeleton for all platforms

go generate ./...

# Run tests first

go vet ./... || exit 1

go test ./... || exit 1

declare -A platforms=([linux]=linux [darwin]=osx [windows]=windows)
declare -A architectures=([386]=i386 [amd64]=amd64)

DESTDIR=dist

mkdir -p $DESTDIR

echo "Building cfn-skeleton"

for platform in ${!platforms[@]}; do
    for architecture in ${!architectures[@]}; do
        echo "... $platform $architecture..."

        name=cfn-skeleton-${platforms[$platform]}-${architectures[$architecture]}

        if [ "$platform" == "windows" ]; then
            name=${name}.exe
        fi

        GOOS=$platform GOARCH=$architecture go build -o $DESTDIR/$name cmd/cfn-skeleton/*
    done
done

echo "All done."
