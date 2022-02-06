#!/bin/bash

output="$(gofmt -d pkg/)"

if [ "$(echo "${output}" | sed '/^\s*$/d' | wc -l)" -gt 0 ]; then
  echo "gofmt detected unformatted files:"
  echo ""
  echo "${output}"
    exit 1
fi
