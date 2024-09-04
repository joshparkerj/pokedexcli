#! /bin/bash
gofmt -w . && staticcheck && gotags ./*.go > tags && go run .

