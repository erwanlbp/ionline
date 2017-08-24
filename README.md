# IOnline
[![Build Status](https://travis-ci.org/erwanlbp/ionline.svg?branch=master)](https://travis-ci.org/erwanlbp/ionline)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/25ae498b822a422a970d147c36246571)](https://www.codacy.com/app/erwan.lbp/ionline?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=erwanlbp/ionline&amp;utm_campaign=Badge_Grade)

## Protobuf

To generate a Go file from a `.proto` file :
```bash
protoc -I=$GOPATH/src/github.com/erwanlbp/ionline/internal/data/protobuf/ --go_out=$GOPATH/src/github.com/erwanlbp/ionline/internal/data/protobuf/ $GOPATH/src/github.com/erwanlbp/ionline/internal/data/protobuf/<filename>.proto
```