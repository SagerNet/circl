#!/usr/bin/env bash

function remove_unused() {
    git rm -rf --ignore-unmatch \
      .github \
      *_test.go **/*_test.go **/test **/testdata \
      mock_*.go **/mock_*.go
}

remove_unused
remove_unused

go mod tidy
