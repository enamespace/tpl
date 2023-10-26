#!/usr/bin/env bash

LC_ALL=C

local_branch="$(git rev-parse --abbrev-ref HEAD)"

valid_branch_regex="^(master|develop)$|(feature|release|hotfix)\/[a-z0-9._-]+$|^HEAD$"

message="wrong branch name"

if [[ ! $local_branch =~ $valid_branch_regex ]]; then
  echo "$message"
  exit 1
fi
exit 0
