#!/bin/bash
echo "Build urmmongo"
go build -ldflags "-s" ../server/main.go
mv main ../dist/urmmongo
cp ../server/config/config.json ../dist/config.json
cp *.sh ../dist
chmod u+x ../dist/*.sh
