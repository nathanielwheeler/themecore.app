#!/bin/bash

cd client
npm install
sass src/styles/global.sass public/global.css
npm run build
cd ..