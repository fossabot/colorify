#!/usr/bin/env bash

set -o errexit
set -o nounset
set -euo pipefail

mkdir -p $HOME/.ssh/ && echo -e "Host github.com\n\tStrictHostKeyChecking no\n" > ~/.ssh/config

TAG=`cat ./.tag`
if [ ! -z "$TAG" ]; then
  echo $TAG
  git tag $TAG
  git push https://${GITHUB_ACCESS_TOKEN}:x-oauth-basic@github.com/${REPO_NAME}/${IMAGE_NAME} --tags
fi
