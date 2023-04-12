#!/bin/bash
# To install cover run
# go get golang.org/x/tools/cmd/cover
go test -coverprofile=c.out
go tool cover -html=c.out -o c.html