#!/bin/bash
# GitHub CLI api
# https://cli.github.com/manual/gh_api

echo -n "Total go repositories: "
gh api \
	-H "Accept: application/vnd.github+json" \
	-H "X-GitHub-Api-Version: 2022-11-28" \
	"/search/repositories?q=language:go" | jq -r .total_count

echo -n "Total go repositories with more than 100 stars: "
gh api \
	-H "Accept: application/vnd.github+json" \
	-H "X-GitHub-Api-Version: 2022-11-28" \
	"/search/repositories?q=stars:>100+language:go" | jq -r .total_count
