#!/bin/bash

# Get the full go repo url
REPO=$(pwd |  rev | cut -d'/' -f-3 | rev)

# Get the name of the app
APP="${PWD##*/}"

# Get this tag as the version
VERSION=$(git describe --abbrev=0 --tags)
# Ensure a fresh build folder
rm -rf build && mkdir build
# Compile
xgo \
  -dest build/ \
  --targets=darwin-10.11/amd64,freebsd/386,freebsd/amd64,freebsd/arm,linux/386,linux/amd64,linux/arm \
  --ldflags "-X main.version=${VERSION}" \
  ../$APP
# Package
cd build
# For each compiled binary, we're repackaging it in a zip with the architecture name, and
# renaming the binary to the app name
for FILE in $(ls .); do
  mv $FILE $APP
  tar cvzf ${FILE}.tgz $APP
  rm -f $APP
done
cd ..