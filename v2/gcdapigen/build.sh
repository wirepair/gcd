#!/bin/sh

go build
./gcdapigen protocol.json
cp -R output/. ../gcdapi
cd ../ && go build && go install
