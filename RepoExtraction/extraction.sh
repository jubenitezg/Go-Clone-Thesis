#!/bin/bash
set -e
source helpers/function-extraction.sh

# Enable debug mode
set -x

cwd=$(pwd)
EXTRACT_EXE="gfunc"
TMP_DIR="/tmp/temp-repo"

setup() {
    build_function_extractor "${EXTRACT_EXE}"
    rm -rf "${TMP_DIR}"
}

main() {
    setup
    pushd "${cwd}/../GitHubMiner/output"
        jq -r '.[:10] | .[] | "\(.name),\(.owner),\(.url)"' "go-repositories.json" | while IFS= read -r line; do
            IFS=',' read -r name owner url <<< "$line"
            git clone "${url}" "${TMP_DIR}"
            "${cwd}/${EXTRACT_EXE}" --project-path "${TMP_DIR}" > "${cwd}/output/${owner}-${name}.json"
            rm -rf "${TMP_DIR}"
        done
    popd
}

main
