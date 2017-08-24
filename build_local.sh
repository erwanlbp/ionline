#!/bin/sh

shortTestMode="-short"
case "-long" in
	"$1")
		shortTestMode="" ;;
esac

START=$(date +%s);

go fmt $(go list ./... | grep -v /vendor/)

goimports -l -w $(find . -name \*.go -print | grep -v /vendor/)

echo "Check Style..."
# Install golint
go get -u github.com/golang/lint/golint

# Check with golint
GOLINT=$(golint $(go list ./... | grep -v /vendor/) 2>&1)
if [ -n "$GOLINT" ]
then
  echo "Non-standard linting in:" >&2
  echo $GOLINT >&2
  echo "FAILED"
  exit 1
fi

# Check with go vet
GOVET=$(go vet $(go list ./... | grep -v /vendor/) 2>&1)
if [ -n "$GOVET" ]
then
  echo "Non-standard constructs in:" >&2
  echo $GOVET >&2
  echo "FAILED"
  exit 1
fi

echo "Building..."
CGO_ENABLED=0 go build
if [ $? -ne 0 ]; then
  echo "FAILED"
  exit 1
fi

echo "Testing..."
go test -cover $shortTestMode $(go list ./... | grep -v /vendor/)
if [ $? -ne 0 ]; then
  echo "FAILED"
  exit 1
fi

END=$(date +%s);
echo $((END-START)) | awk '{print "Build took "int($1/60)"min "int($1%60)"sec"}'
echo "SUCCESS"