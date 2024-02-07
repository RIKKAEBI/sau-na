#!/bin/sh

# web component build
cd components/
npm install
npm run build
npm run sb:build
cd ..

# go project build
go build -tags netgo -ldflags '-s -w' -o app

