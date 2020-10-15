#!/bin/bash

cd $GOSRC/themecore.app

source scripts/getenv.sh

# Start client
cd client
rollup -c -w &

cd ..

# Start server
fresh