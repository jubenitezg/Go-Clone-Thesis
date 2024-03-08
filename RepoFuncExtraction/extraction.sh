#!/bin/bash
set -e
source helpers/function-extraction.sh

# Enable debug mode
# set -x

cwd=$(pwd)
EXTRACT_EXE="gfunc"
TMP_DIR="/tmp/temp-repo"

print_usage() {
  printf "Usage: -o for output directory (needed)\n"
}

setup() {
    build_function_extractor "${EXTRACT_EXE}"
    rm -rf "${TMP_DIR}"
}

main() {
    setup
    pushd "${cwd}/../GitHubMiner/output"
        i=15800
        jq -r '.[15800:15865] | .[] | "\(.name),\(.owner),\(.url)"' "go-repositories.json" | while IFS= read -r line; do
            IFS=',' read -r name owner url <<< "$line"
            echo "Cloning repository: ${owner}/${name}"
            git clone "${url}" "${TMP_DIR}"
            echo "Extracting functions from repository: ${owner}/${name}"
            "${cwd}/${EXTRACT_EXE}" --project-path "${TMP_DIR}" --single-line > "${OUTPUT_DIR}/${owner}-${name}.json"
            rm -rf "${TMP_DIR}"
            echo "Processed repository: ${i}"
            echo "Processed repository: ${i}" > "${cwd}/current.log"
            i=$((i+1))
        done
    popd
}

while getopts 'o:' flag; do
  case "${flag}" in
    o) OUTPUT_DIR="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done
if [ -z "${OUTPUT_DIR}" ]; then
    print_usage
    exit 1
fi

main
