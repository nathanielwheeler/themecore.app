#!/bin/bash

cd $GOSRC/themecore.app

# Start client
cd client
rollup -c -w &

cd ..

# Start server
fresh