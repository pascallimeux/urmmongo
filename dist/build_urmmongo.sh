#!/bin/bash
echo "Build urmmongo"
go build -ldflags "-s" $GOPATH/src/github.com/pascallimeux/urmmongo/server/main.go
if [ ! -d "$GOPATH/src/github.com/pascallimeux/urmmongo/dist" ]; then
  mkdir $GOPATH/src/github.com/pascallimeux/urmmongo/dist
fi
mv main $GOPATH/src/github.com/pascallimeux/urmmongo/dist/urmmongo
cp $GOPATH/src/github.com/pascallimeux/urmmongo/server/config/config.json $GOPATH/src/github.com/pascallimeux/urmmongo/dist/config.json
cp *.sh $GOPATH/src/github.com/pascallimeux/urmmongo/dist
chmod u+x $GOPATH/src/github.com/pascallimeux/urmmongo/dist/*.sh
