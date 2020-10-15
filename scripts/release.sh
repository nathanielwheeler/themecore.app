#!/bin/bash
if [[ ! -e main.go ]]; then
	echo "Invalid directory"
	exit 1
fi

# Ensure branches are clean
echo "Checking working tree..."
if [[ -n $(git status -s) || $(git diff --stat) != '' ]]; then
	# Abort if they aren't
	echo "Working tree is dirty.  Release aborted."
	exit 1
fi
echo "Working tree is clean."

# Test locally
testResult=$(go test .)
if [[ $testResult == *"FAIL"* ]]; then
	echo "Tests failed.  Aborting release."
	exit 1
fi

