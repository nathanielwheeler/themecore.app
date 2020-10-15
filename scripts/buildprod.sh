#!/bin/bash
# To be executed in production environment

echo "Starting Production Build"
srcDir=$HOME/go/src/themecore.app
prodDir=$HOME/app

# Update dependencies
cd $srcDir
echo "Updating server dependencies..."
git pull origin main
/usr/local/go/bin/go get -u ./...
go test -i .
echo "Updating client dependencies..."
source $srcDir/client/clientbuild.sh

# Push changes to github
echo "Pushing changes to origin..."
git add .
git commit -m "automatic dependency update"
git push origin main

# Test changes in production
testResult=$(go test .)
if [[ $testResult == *"FAIL"* ]]; then
	echo "Production tests failed.  Aborting release."
	exit 1
fi

# Delete local binary and client if it exists
cd $prodDir
echo "Removing previous build..."
rm server
rm -rf client

# Build
echo "Building server..."
/usr/local/go/bin/go build -o ./server $srcDir/*.go
echo "Building client..."
cp -R $srcDir/client .
cp $srcDir/Caddyfile .

# Restart server
echo "Restarting server..."
service themecore.app restart
service caddy restart

echo "Finished!"
