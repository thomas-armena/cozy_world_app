#!/bin/sh

protoc --proto_path=protos/ --go_out=server/cozyworld/protos/ --go_opt=paths=source_relative protos/*.proto
protoc --proto_path=protos/ --dart_out=lib/protos/ protos/*.proto

