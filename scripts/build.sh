#!/bin/sh
cd /go/src/app
go build -o /dist/app app
cp /dist/app /go/bin/labellab
