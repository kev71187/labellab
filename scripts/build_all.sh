#!/bin/sh
cd /go/src/app
go build -o /dist/app app
cp /dist/app /go/bin/labellab
env GOOS=linux GOARCH=amd64 go build -o /dist/llab.ubuntu main.go
# env GOOS=linux GOARCH=arm go build -o llab.ubuntu main.go
env GOOS=darwin GOARCH=amd64 go build -o /dist/llab.mac main.go

