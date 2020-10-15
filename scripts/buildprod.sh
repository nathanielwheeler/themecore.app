#!/bin/bash
# To be executed in production environment
echo "Starting Production Build"
cd $HOME/app

# Delete local binary and client if it exists
echo "Removing previous build..."
rm themecore.app
rm -rf client

# Update dependencies
cd source
echo "Updating dependencies..."
git pull origin main
/usr/local/go/bin/go get -u ./...

# Push changes to github
echo "Pushing changes to origin..."
git add .
git commit -m "automatic dependency update"
git push origin main
cd ..

# Build
echo "Building server..."
/usr/local/go/bin/go build -o ./themecore.app ./source/main.go
echo "Building client..."
cp -R ./source/client .
cp ./source/Caddyfile .

# Restart server
echo "Restarting server..."
service themecore.app restart
service caddy restart

echo "Finished!"
