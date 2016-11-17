#!/usr/bin/env bash

# This script gathers Go test results for uploading to codecov.io

set -e
echo "" > coverage.txt

for d in $(find ./* -maxdepth 10 -type d -not -path "*vendor*"); do
	if ls $d/*.go &> /dev/null; then
		go test -coverprofile=profile.out -covermode=atomic $d
		if [ -f profile.out ]; then
			cat profile.out >> coverage.txt
			rm profile.out
		fi
	fi
done
