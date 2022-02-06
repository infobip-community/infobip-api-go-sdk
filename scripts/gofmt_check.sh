#!/bin/bash

output="$(gofmt -d "$(go list ./... | sed 's/infobip-go-client\///g')")"

if [ "$(echo "${output}" | sed '/^\s*$/d' | wc -l)" -gt 0 ]; then
  echo "gofmt detected unformatted files:"
  echo ""
  echo "${output}"
    exit 1
fi
