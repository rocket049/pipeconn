#!/bin/sh

go build pipe-server.go
go build pipe-client.go
./pipe-client
