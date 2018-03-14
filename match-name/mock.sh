#!/bin/sh

set -x

mkdir -p mock

mkdir -p mock/name-parser
mockgen -package "mock_np" -destination "./mock/name-parser/name-parser.go" "github.com/uberblah/go-demo/name-parser" "NameParser"
