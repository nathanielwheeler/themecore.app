#!/bin/bash
if [[ ! -e main.go ]]; then
	echo "Invalid directory"
	exit 1
fi

# Start client
cd client
rollup -c -w &

cd ..

# Start server
fresh