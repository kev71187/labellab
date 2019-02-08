#!/bin/sh
cd /go/src/app
go build -ldflags "-X app/config.Env=local" -o /dist/app app
cp /dist/app /go/bin/labellab
