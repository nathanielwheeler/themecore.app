#!/bin/bash
if [[ ! -e main.go ]]; then
	echo "Invalid directory.  Process aborted"
	exit 1
fi

# Ensure branches are clean
echo "Checking working tree..."
if [[ -n $(git status -s) || $(git diff --stat) != '' ]]; then
	# Abort if they aren't
	echo "Working tree is dirty.  Process aborted."
	exit 1
fi
echo "Working tree is clean."

# Test locally
echo "Testing package..."
testResult=$(go test .)
if [[ $testResult == *"FAIL"* ]]; then
	echo "Tests failed.  Process aborted."
	exit 1
fi

echo "All tests passed!"