#!/usr/bin/env bash

set -o errexit
set -o nounset
set -euo pipefail

mkdir -p $HOME/.ssh/ && echo -e "Host github.com\n\tStrictHostKeyChecking no\n" > ~/.ssh/config

LAST_COMMIT=`git log -1 --pretty=%B`
touch ./.tag
if VERSION=`git describe --abbrev=0 --tags` && [ ! -z "`git diff $VERSION`" -o -z "$VERSION" ]; then
  VERSION=${VERSION:-'0.0.0'}
  MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
  MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
  PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"
  if echo $LAST_COMMIT | grep "\[\(major\|MAJOR\)\]" > /dev/null; then
    MAJOR=$((MAJOR+1))
    echo "$MAJOR.0.0" > ./.tag
  elif echo $LAST_COMMIT | grep "\[\(minor\|MINOR\)\]" > /dev/null; then
    MINOR=$((MINOR+1))
    echo "$MAJOR.$MINOR.0" > ./.tag
  elif echo $LAST_COMMIT | grep "\[\(patch\|PATCH\)\]" > /dev/null; then
    PATCH=$((PATCH+1))
    echo "$MAJOR.$MINOR.$PATCH" > ./.tag
  fi
else
  if echo $LAST_COMMIT | grep "\[\(major\|MAJOR\)\]" > /dev/null; then
    echo "v1.0.0" > ./.tag
  elif echo $LAST_COMMIT | grep "\[\(minor\|MINOR\)\]" > /dev/null; then
    echo "v0.1.0" > ./.tag
  elif echo $LAST_COMMIT | grep "\[\(patch\|PATCH\)\]" > /dev/null; then
    echo "v0.0.1" > ./.tag
  fi
fi
