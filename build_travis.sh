#!/bin/sh

echo "Check Go fmt..."
GOFMT=$(go fmt $(go list ./... | grep -v /vendor/))
if [ -n "$GOFMT" ]
then
  echo "Non-standard formatting in:"
  echo "$GOFMT"
  exit 1
fi

echo "Check goimports..."
# Install goimports
go get golang.org/x/tools/cmd/goimports

GOIMP=$(goimports -l -w $(find . -name \*.go -print | grep -v /vendor/))
if [ -n "$GOIMP" ]
then
  echo "Non-standard imports format in:"
  echo "$GOIMP"
  exit 1
fi

/bin/bash build_local.sh
if [ $? -ne 0 ]; then
  exit 1
fi

echo "Send cover to goveralls..."
go get github.com/mattn/goveralls
goveralls -coverprofile=acc.coverprofile -service=travis-ci -repotoken $COVERALLS_TOKEN
if [ $? -ne 0 ]; then
  echo "FAILED"
  exit 1
fi
