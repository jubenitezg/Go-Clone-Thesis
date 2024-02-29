#!/bin/bash
set -e
# Enable to debug
set -x

cwd=$(pwd)

build_function_extractor() {
    pushd "${cwd}/../GoFuncExtractor"
        go build -o "${1}"
        cp "${1}" "${cwd}"
    popd
}
