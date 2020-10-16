#!/bin/bash
if [[ ! -e main.go ]]; then
	echo "Invalid directory"
	exit 1
fi

# Start client
cd client
sass --watch src/styles:public &
rollup -c -w &

cd ..

# Start server
fresh