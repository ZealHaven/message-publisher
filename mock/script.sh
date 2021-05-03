#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# lib
mockgen -package=mock -destination=mock/mock_response_writer.go -source=/src/net/http/server.go