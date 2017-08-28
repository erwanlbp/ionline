#!/bin/bash

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
GOLINT=$(golint $(go list ./... | grep -v /vendor/))
if [ -n "$GOLINT" ]
then
  echo "Non-standard linting in:" >&2
  echo "$GOLINT"
  echo "FAILED"
  exit 1
fi

# Check with go vet
GOVET=$(go vet $(go list ./... | grep -v /vendor/))
if [ -n "$GOVET" ]
then
  echo "Non-standard constructs in:" >&2
  echo "$GOVET"
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
echo "-- Testing cmd ..."
go test -cover $shortTestMode ./cmd/...
if [ $? -ne 0 ]; then
  echo "FAILED"
  exit 1
fi

echo "-- Testing internal ..."
echo "mode: set" > acc.coverprofile
for Dir in $(go list ./internal/...);
do
    returnval=$(go test -coverprofile=profile.out $Dir $shortTestMode -args -public $GOPATH/src/github.com/erwanlbp/ionline/internal/public/ -firebase-auth IONLINE_TEST_SECRET_FIREBASE -log stdout)
    echo "$returnval"
    if [[ ${returnval} != *FAIL* ]]
    then
        if [ -f profile.out ]
        then
            cat profile.out | grep -v "mode: set" >> acc.coverprofile
        fi
    else
        exit 1
    fi
done

END=$(date +%s);
echo $((END-START)) | awk '{print "Build took "int($1/60)"min "int($1%60)"sec"}'
echo "SUCCESS"