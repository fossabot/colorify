#!/usr/bin/env bash

set -o errexit
set -o nounset
set -euo pipefail

# curl -sL https://git.io/goreleaser | bash
mkdir -p $HOME/.ssh/ && echo -e "Host github.com\n\tStrictHostKeyChecking no\n" > ~/.ssh/config

TAG=`cat ./.tag`
if [ ! -z "$TAG" ]; then
  echo "Create release: ${TAG}"
  curl -H "Authorization: token ${GITHUB_ACCESS_TOKEN}" \
       -X POST \
       -d "{\"tag_name\": \"${TAG}\"}" \
       ${GITHUB_API}repos/${REPO_NAME}/${IMAGE_NAME}/releases
fi
