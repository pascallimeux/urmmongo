#!/bin/bash
echo "Build urmmongo"
go build -ldflags "-s" main.go
mv server ../dist/urmmongo
cp config.json ../dist/config.json
cp ../scripts/*.sh ../dist
