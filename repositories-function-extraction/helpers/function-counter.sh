#!/bin/bash

#set -x

print_usage() {
  printf "Usage: -p for input directory (needed)\n"
}

main() {
  cwd=$(pwd)
  pushd "$INPUT_DIR" || exit
  total_count=0

  for file in *.json; do
    cp "$file" "/tmp/$file"
    count=$(jq length "/tmp/$file")
    total_count=$((total_count + count))
    echo "file: $file"
    echo "count of functions in $file: $count"
    echo "count of functions in $file: $count" >>"${cwd}/count.log"
    rm "/tmp/$file"
  done

  echo "Total count of functions: $total_count"

  echo "Total count: $total_count" >>"${cwd}/count.log"
  popd || exit
}

while getopts 'p:' flag; do
  case "${flag}" in
  p) INPUT_DIR="${OPTARG}" ;;
  *)
    print_usage
    exit 1
    ;;
  esac
done
if [ -z "${INPUT_DIR}" ]; then
  print_usage
  exit 1
fi

main
