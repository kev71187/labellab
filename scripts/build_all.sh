#!/bin/sh
export APP_ENV=production
cd /go/src/app
go build -ldflags "-X app/config.Env=production" -o /dist/app app
cp /dist/app /go/bin/labellab
env GOOS=linux GOARCH=amd64 go build -ldflags "-X app/config.Env=production" -o /dist/llab.ubuntu main.go
# env GOOS=linux GOARCH=arm go build -o llab.ubuntu main.go
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X app/config.Env=production" -o /dist/llab.mac main.go

