#!/usr/bin/env bash

echo Testing bad prompt
cat << EOF | go run ./cmd/windchill
46
13
10
EOF
