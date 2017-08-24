#!/bin/sh

echo "Check Go fmt..."
GOFMT=$(go fmt $(go list ./... | grep -v /vendor/) 2>&1)
if [ -n "$GOFMT" ]
then
  echo "Non-standard formatting in:" >&2
  echo $GOFMT >&2
  exit 1
fi

echo "Check goimports..."
GOIMP=$(goimports -l -w $(find . -name \*.go -print | grep -v /vendor/) 2>&1)
if [ -n "$GOIMP" ]
then
  echo "Non-standard imports format in:" >&2
  echo $GOIMP >&2
  exit 1
fi

/bin/sh build_local.sh